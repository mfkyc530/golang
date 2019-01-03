package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add", TestAddUser)
	t.Run("Get", TestGetUser)
	t.Run("Del", TestDeleteUser)
	t.Run("Reget", TestRegetUser)
}

func TestAddUser(t *testing.T){
	err := AddUserCredential("wch","123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUser(t *testing.T){
	pwd, err := GetUserCredential("wch")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("wch","123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func TestRegetUser(t *testing.T){
	pwd, err := GetUserCredential("wch")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != ""{
		t.Errorf("Deleting user test failed")
	}
}


func TestVideoWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser", TestAddUser)
	t.Run("AddVideo", TestAddVideoInfo)
	t.Run("GetVideo", TestGetVideoInfo)
	t.Run("DelVideo", TestDeleteVideoInfo)
	t.Run("RegetVideo", TestRegetVideoInfo)
}

func TestAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1,"my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func TestGetVideoInfo(t *testing.T) {
	_ ,err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func TestDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func TestRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T){
	clearTables()
	t.Run("AddUser", TestAddUser)
	t.Run("AddComments", TestAddComments)
	t.Run("ListComments", TestListComments)
}

func TestAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComments(vid, content, aid)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func TestListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil{
		t.Errorf("Error of TestListComments: %v", err)
	}

	for i, ele := range res{
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}

