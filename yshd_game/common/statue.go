package common

//房间状态类型
const (
	ROOM_NONE   = iota
	ROOM_ONLIVE = 1
	//	ROOM_PRE      = 2
	ROOM_FINISH   = 3
	ROOM_MODIFY   = 4
	ROOM_PLAYBACK = 5
	ROOM_MULTIPLE = 6
	//ROOM_READY    = 7
	ROOM_RESTART = 7
	ROOM_READY   = 8

	ROOM_PRE_V2 = 9
)

//主播状态
const (
	USER_STATUE_LEAVE  = iota
	USER_STATUE_LIVE   = 1
	USER_STATUE_RELIVE = 3
	USER_STATUE_SEE    = 4
)

//提现状态
const (
	CASH_BANK_STATUE_NONE = iota
	CASH_BANK_STATUE_VET
	CASH_BANK_STATUE_VET_FAILED  //审核失败2
	CASH_BANK_STATUE_SUCCESS     //审核成功3
	CASH_PAY_SUCCESS             //打款成功4
	CASH_PAY_FAIL                //打款失败5
	CASH_PAY_RICE_RETURN         //6
	CASH_PAY_NOTENOUGH           //7
	CASH_PAY_TENCENT_SYSTEMERROR //8
	//CASH_BANK_STATUE_REFUSE
)

//月亮商城订单状态 1审核中 2拒绝 3线上发货成功 4线下发货 5线上打款失败
const (
	SCORE_EXCHANGE_STATUS_NONE            = iota
	SCORE_EXCHANGE_STATUS_VET             = 1 //审核中
	SCORE_EXCHANGE_STATUS_FAILED          = 2 //拒绝
	SCORE_EXCHANGE_STATUS_ONLINE_SUCCESS  = 3 //线上打款成功
	SCORE_EXCHANGE_STATUS_OFFLINE_SUCCESS = 4 //线下发货
	SCORE_EXCHANGE_STATUS_ONLINE_FAILED   = 5 //线上打款失败
)

const (
	MOON_CASH_STATUS_VET     = 1 //月亮商城提现审核中
	MOON_CASH_STATUS_SUCCESS = 2 //线上发货
	MOON_CASH_STATUS_FAILED  = 3 //线上打款失败
)

//订单支付状态 1预创建2成功3取消
const (
	TRADE_NONE = iota
	TRADE_PRE_CREATE
	TRADE_HAVE_SUCCESS
	TRADE_CANCELED
)

//实名认证状态
const (
	REAL_AUTH_VET = iota
	REAL_AUTH_REFUSE
	REAL_AUTH_OK
)

//封号类型
const (
	FORBID_POWERS_NONE = iota
	FORBID_POWERS_FOREVER
	FORBID_POWERS_TIME
	FORBID_POWER_MAX
)

//房间状态
const (
	MULTIPLE_ROOM_PRE  = iota
	MULTIPLE_ROOM_BUSY = 1
	//MULTIPLE_ROOM_FREE = 2
	MULTIPLE_ROOM_FIN  = 2
	MULTIPLE_ROOM_LOCK = 4
)

//房间类型
const (
	ROOM_TYPE_NOMARL = iota
	ROOM_TYPE_TEST   = 1
)

//关注的状态
const (
	FOCUS_STATUE_NONE = iota
	FOCUS_STATUE_MAIN
)

//交易类型
//1微信2苹果3谷歌4h5支付
const (
	TRADE_SELF           = iota
	TRADE_TYPE_WEIXIN    = 1
	TRADE_TYPE_APPLE     = 2
	TRADE_TYPE_GOOGLE    = 3
	TRADE_TYPE_WEIXIN_H5 = 4
)

const (
	//三方支付交易类型
	THIRD_PAY_UNIONPAY = 11
	THIRD_PAY_ALIPAY   = 12
	THIRD_PAY_WEIXIN   = 13
	//
	ALIPAY_PAY = 29
	HUAWEI_PAY = 30
)
