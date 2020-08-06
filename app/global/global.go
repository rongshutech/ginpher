package global

const (
	/*
		redis
	*/
	RedisKeySmsLogin = "tf:app:sms:login:%s" //登录时验证码
	RedisKeySmsRegis = "tf:app:sms:regis:%s" //注册时验证码
	RedisKeySmsMPwd  = "tf:app:sms:mpwd:%s"  //修改密码验证码
	RedisKeySmsFPwd  = "tf:app:sms:fpwd:%s"  //忘记密码验证码
	RedisKeyToken    = "tf:app:token:%s"     //token
)

const (
	/*
		用户
	*/
	//Registering  = 1 // 注册中
	Register     = 0 // 注册成功
	Cancellation = 4 // 注销

	/*
		登录方式
	*/
	LoginTypePwdM        = 1 // 手机号密码登录
	LoginTypeVerifyCode  = 2 // 手机号验证码登录
	LoginTypeFace        = 3 // 人脸登录
	LoginTypeFingerprint = 4 // 指纹登录

	/*
		redis存储的验证码的类型
	*/
	TypeLogin     = 1 // 登录
	TypeRegister  = 2 // 注册
	TypeModifyPwd = 3 // 修改密码
	TypeForgetPwd = 4 // 忘记密码

	/*
		个人标签
	*/
	AuditNo      = 0 //待审核
	Auditing     = 1 //审核中
	AuditSuccess = 2 //审核成功
	AuditFail    = 3 //审核失败
	AuditDelete  = 4 //删除标签
	Uncertified  = "未认证"
	//ToBeAuditedMsg       = "待审核"
	ToBeUploadedMsg      = "未上传"
	InAuditMsg           = "智能分析中"
	AuditSuccessMsg      = "已更新"
	CertificationSuccess = "认证成功"
	AuditFailStr         = "审核失败"
	Certified            = "已认证"
	UnWrite              = "未填写"

	LabelIdName    = 1 // 姓名
	LabelMobile    = 2 // 手机号
	LabelFirm      = 3 // 公司/职位
	LabelEducation = 4 // 学历
	LabelImage     = 5 // 个人照片
	LabelVideo     = 6 // 个人展示视频
	LabelPurpose   = 7 // 默认去电意向
	LabelThrough   = 8 // 通行大数据行程卡

	/*
		通话
	*/
	CallTypeTaoFu        = 1 // 桃福通话
	CallTypeTaoFuPrivacy = 2 // 桃福隐私通话
	CallTypeCommon       = 3 // 普通通话

	StatusOut     = 1 // 已拨出
	StatusReceive = 2 // 收到
	StatusAnswer  = 3 // 接听
	StatusHangUp  = 4 // 挂断
	StatusRefuse  = 5 // 拒接

	BellDuration = 60 // 响铃60s后自动挂断

	/*
		订单状态
	*/
	OrderStatus_NoPay    = 1 // 未支付
	OrderStatus_Pay      = 2 // 已支付
	OrderStatus_Success  = 3 // 支付成功
	OrderStatus_Fail     = 4 // 支付失败
	OrderStatus_OverTime = 5 // 支付超时

	/*
		支付类型  //TODO
	*/
	PayType_Wechat = 1 // 微信支付

	/*
		活动
	*/
	ActivityType_New        = 1 // 新人活动
	ActivityType_Five       = 2 // 5元活动
	ActivityType_Ten        = 3 // 10元活动
	ActivityType_TwentyFive = 4 // 25元活动

	/*
		兑换码
	*/
	ExchangeCode_NoActive = 0 // 未激活
	ExchangeCode_Active   = 1 // 激活
	ExchangeCode_Used     = 2 // 已使用

	/*
		邀请
	*/
	InviteType_Register = 1 // 注册邀请
	InviteType_May      = 2 // 购买邀请

	/*
		mysql table name
	*/
	TableUsers          = "users"           // 用户表
	TableLabels         = "labels"          // 标签表
	TablePersonalLabels = "personal_labels" // 个人标签表
	TableCallRecords    = "call_records"    // 通话记录表
	TableOrder          = "order"           // 订单表
	TableActivities     = "activities"      // 活动表
	TableGoods          = "goods"           //商品表
	TableExchangeCodes  = "exchange_codes"  // 兑换码表
	TableInvites        = "invites"         // 邀请表
	TableAgreements     = "agreements"      // 协议表
	TableLoginLogs      = "login_logs"      // 登录日志表
	TableQuestions      = "questions"       // 客服问题表
	TableVersions       = "versions"        // 版本表
)
