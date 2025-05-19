package libbox

const (
	CommandLog int32 = iota
	CommandStatus
	CommandServiceReload
	CommandServiceClose
	CommandCloseConnections
	CommandGroup
	CommandSelectOutbound
	CommandURLTest
	CommandGroupExpand
	CommandClashMode
	CommandSetClashMode
	CommandGetSystemProxyStatus  
	CommandSetSystemProxyEnabled
	CommandGroupInfoOnly  //一定要排在这个位置，13，这样才能和客户端的startcommand的参数13一一对应
	CommandConnections
	CommandCloseConnection
	CommandGetDeprecatedNotes

)
