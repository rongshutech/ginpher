package api

type Code int

const (
	// Success 成功
	Success Code = 0

	// BadRequest 参数错误相关
	BadRequest                  Code = 100000
	LackParams                  Code = 100001
	IllegalParams               Code = 100002
	RPCError                    Code = 100003
	IllegalRequest              Code = 100004
	CodeAcquisitionFailed       Code = 100005
	VerificationCodeFailInSend  Code = 100006
	VerificationCodeError       Code = 100007
	VerificationCodeExpired     Code = 100008
	VerificationCodeOversending Code = 100009
	DataAlreadyExists           Code = 100010
	LinkExpired                 Code = 100011
	OffendingWord               Code = 100012
	VerificationCodeSendError   Code = 100013
	EncryptionFailed            Code = 100014
	PacketDecryptionFailed      Code = 100015
	RealNameFailure             Code = 100016
	UploadFailure               Code = 100017
	DifferentPersonalInfo       Code = 100018
	TokenExpired                Code = 100019

	// Unauthorized 权限错误相关
	Unauthorized          Code = 100100
	NeedLoginFirst        Code = 100101
	NoRightToLogin        Code = 100102
	NeedHigherAuthority   Code = 100103
	NoRightToVisit        Code = 100104
	NoRightToDelele       Code = 100105
	NoRightToCreate       Code = 100106
	NoRightToUpdate       Code = 100107
	NoRightToBuy          Code = 100108
	NoRightThreePartyApi  Code = 100109
	StructuralAbandonment Code = 100110
	DidNotSubscribe       Code = 100111
	NoRegister            Code = 100112
	Registered            Code = 100113
	NoRealName            Code = 100114

	// Forbidden 禁止相关操作
	Maintenance          Code = 100301
	ResourcesUnavailable Code = 100302
	SystemBusy           Code = 100303
	LoginTimeout         Code = 100304
	LoginFailure         Code = 100305
	CreateFailed         Code = 100306
	QueryFailed          Code = 100307
	DeleteFailed         Code = 100308
	FrequentlyOpration   Code = 100309
	ChannelUrlTimeOut    Code = 100310
	ChannelUrlFail       Code = 100311
	PasswordFalse        Code = 100312
	PasswordNoSame       Code = 100313
	OldPwdFalse          Code = 100314
	ModifyPwdFail        Code = 100315
	UsedNumFail          Code = 100316

	// NotFound 未找到
	NotFound    Code = 100400
	ApiNotFound Code = 100401
	NoData      Code = 100402
	NoUser      Code = 100403

	Forbidden Code = 100500

	OtherError Code = 999
)

var (
	statusText = map[Code]string{
		Success:                     "成功",
		BadRequest:                  "错误的请求",
		LackParams:                  "缺失必选参数",
		IllegalParams:               "非法参数",
		RPCError:                    "RPC错误",
		IllegalRequest:              "非法请求",
		CodeAcquisitionFailed:       "Code获取失败",
		LoginTimeout:                "登录超时，请重新登录",
		LoginFailure:                "登录失败",
		CreateFailed:                "创建失败",
		QueryFailed:                 "数据获取失败",
		DeleteFailed:                "删除失败",
		FrequentlyOpration:          "操作过于频繁",
		VerificationCodeFailInSend:  "验证码超限发送",
		VerificationCodeError:       "验证码错误",
		VerificationCodeExpired:     "验证码失效",
		VerificationCodeOversending: "验证码发送频繁",
		DataAlreadyExists:           "用户已存在",
		LinkExpired:                 "链接失效",
		OffendingWord:               "包含敏感词",
		EncryptionFailed:            "数据加密失败",
		PacketDecryptionFailed:      "数据包解密失败",
		RealNameFailure:             "实名认证失败",
		UploadFailure:               "上传失败请重新上传",
		DifferentPersonalInfo:       "个人信息比对失败",
		TokenExpired:                "Token过期",

		Unauthorized:          "未授权",
		NeedLoginFirst:        "未登录",
		NoRightToLogin:        "没有权限登录",
		NeedHigherAuthority:   "该资源需要拥有更高级的授权",
		NoRightToVisit:        "无权访问",
		NoRightToDelele:       "无权删除",
		NoRightToCreate:       "无权新增",
		NoRightToUpdate:       "无权修改",
		NoRightToBuy:          "无权购买",
		NoRightThreePartyApi:  "第三方应用访问api接口权限受限制",
		StructuralAbandonment: "该接口已经废弃",
		DidNotSubscribe:       "未关注",
		NoRegister:            "未注册",
		Registered:            "已注册",
		NoRealName:            "未实名",

		Forbidden:                 "拒绝访问",
		Maintenance:               "服务器维护中",
		ResourcesUnavailable:      "服务端资源不可用",
		SystemBusy:                "任务过多，系统繁忙",
		NotFound:                  "拒绝访问",
		ApiNotFound:               "接口不存在",
		NoData:                    "没有相关数据",
		NoUser:                    "用户不存在",
		VerificationCodeSendError: "验证码发送失败",

		OtherError:        "其他错误",
		ChannelUrlTimeOut: "链接已超时！",
		ChannelUrlFail:    "链接无效",
		PasswordFalse:     "密码错误",
		PasswordNoSame:    "密码不一致",
		OldPwdFalse:       "原密码错误",
		ModifyPwdFail:     "修改密码失败",
		UsedNumFail:       "次数已达上限",
	}
)

func (c Code) Msg() string {
	return statusText[c]
}
