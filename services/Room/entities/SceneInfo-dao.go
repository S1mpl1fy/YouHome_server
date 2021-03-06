package entities

import "encoding/json"

func AddRoom(roomName , userId,roomUrl string) error {
	sql := "INSERT INTO room (roomName, userId,url) VALUES (?, ?,?)"
	_, err := mydb.Exec(sql, roomName, userId,roomUrl)
	return err
}

func GetRooms(Userid string) string {
	sql := "SELECT roomId,roomName,url FROM room WHERE userId = ?"
	rows, err := mydb.Query(sql, Userid)
	checkErr(err)

	type Fat struct {
		RoomId string `json:"roomId"`
		RoomName string `json:"roomName"`
		Url string `json:"url"`
	}
	var temp Fat
	result := make([]Fat, 0)
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&temp.RoomId,&temp.RoomName,&temp.Url)
		checkErr(err)
		result = append(result,temp)
	}
	re , err := json.Marshal(result)
	checkErr(err)
	return string(re)
}

func GetRoomId(roomName string, userId string) string{
	sql := "SELECT roomId FROM room WHERE roomName = ? and userId = ? ORDER BY roomId"
	rows, err := mydb.Query(sql, roomName,userId)
	checkErr(err)
	var roomId string
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&roomId)
	}
	return roomId
}
func GetRoomName(roomId string)string{
	sql := "select roomName from room where roomId = ?"
	rows, err := mydb.Query(sql, roomId)
	checkErr(err)
	var roomName string
	if rows == nil{
		return ""
	}
	for rows.Next(){
		rows.Scan(&roomName)
	}
	return roomName
}
func GetRoomUrl(roomId string)string{
	sql := "select url from room where roomId = ?"
	rows, err := mydb.Query(sql,roomId)
	checkErr(err)
	var roomUrl string
	if rows == nil{
		return ""
	}
	for rows.Next(){
		rows.Scan(&roomUrl)
	}
	return roomUrl;
}
func UpdateRoomName(roomName string,roomId string)string{
	sql := "UPDATE room SET roomName = ? WHERE roomId = ?"
	_ ,err := mydb.Exec(sql,roomName,roomId)
	checkErr(err)
	return roomName
}
func UpdateRoomUrl(roomUrl,roomId string)bool{
	sql := "update room set url = ? where roomId = ?"
	_, err := mydb.Exec(sql,roomUrl,roomId)
	if err != nil{
		return false
	}
	return true
}
func DeleteRoom(id string) error {
	sql := "delete FROM room where roomId=? "
	_, err := mydb.Exec(sql, id)
	return err
}