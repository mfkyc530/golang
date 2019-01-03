package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"log"
	"time"
	"vido_server/api/defs"
)

func AddUserCredential(loginName string, pwd string) error{
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) VALUES (?, ?)")
	if err != nil{
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error){
	stmOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil{
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "", err
	}
	defer stmOut.Close()

	return pwd, nil
}

func DeleteUser(loginName, pwd string)error{
	stmDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}

	_, err = stmDel.Exec(loginName, pwd)
	if err != nil{
		return  err

	}
	defer stmDel.Close()
	return nil
}


func AddNewVideo(aid int, name string) (*defs.VideoInfo, error){
	// create uuid
	u1 := uuid.Must(uuid.NewV4())
	var vid string
	for i := 0; i < len(u1); i++ {
		if u1[i] == 0 {
			vid = string(u1[0:i])
		}
	}

	t := time.Now()
	ctime :=  t.Format("Jan 02 2006, 15:04:05")
	stmIns, err := dbConn.Prepare(`INSERT INTO video_info(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)
	if err != nil{
		return nil, err
	}

	_, err = stmIns.Exec(vid, aid, name, ctime)
	if err != nil{
		return nil,err
	}

	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtine:ctime}

	defer stmIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error){
	stmOut, err := dbConn.Prepare("SELECT id,author_id, name, display_ctime FROM video_info WHERE id = ?")

	var aid int
	var dct string
	var name string
	err = stmOut.QueryRow(vid).Scan(&vid, &aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows{
		return nil, nil
	}

	defer stmOut.Close()
	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtine:dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error{
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return  err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil{
		return  err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid, content string, aid int) error{
	u1 := uuid.Must(uuid.NewV4())
	var id string
	for i := 0; i < len(u1); i++ {
		if u1[i] == 0 {
			id = string(u1[0:i])
		}
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return  err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error){
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name,comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next(){
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id:id, VideoId:vid, Author:name, Content:content}
		res = append(res, c)
	}

	defer stmtOut.Close()
	return res, nil
}
