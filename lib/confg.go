package lib

type Iconfig interface {
	//提供接口 api账号和 口令的接口
	GetApiUserAndPaw() (string, string)

	CompanyID() int64
}
