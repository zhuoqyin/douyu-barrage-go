package douyu

import (
	"strconv"
	"time"
)

func LoginReq(roomID string) []byte {
	return []byte("type@=loginreq/roomid@=" + roomID + "/")
}

func JoinGroupReq(roomID string) []byte {
	return []byte("type@=joingroup/rid@=" + roomID + "/gid@=-9999/")
}

func KeepaliveReq() []byte {
	return []byte("type@=keeplive/tick@=" + strconv.FormatInt(time.Now().Unix(), 10) + "/")
}
