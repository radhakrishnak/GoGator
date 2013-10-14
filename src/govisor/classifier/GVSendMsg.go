package classifier

import (

)
type GVSendMsg interface{
	GetConnectionName() string
	GetName() string
	SendMsg(from GVSendMsg)
	DropMsg(from GVSendMsg)
}

