package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamempmpsessionsendmessagetogameAPIRequest 发送消息给游戏 API请求
// alibaba.cgame.mp.mpsession.sendmessagetogame
//
// 发送消息给游戏
type AlibabacgamempmpsessionsendmessagetogameAPIRequest struct {
	model.Params
	// 设备Id， 可以传空值
	_deviceId string
	// 调用者的userid，若有userAccessToken,可以传空值
	_userId string
	// 目前没有使用，可以传空值
	_userToken string
	// 登录后获取的accessToken
	_userAccessToken string
	// ecs上的实例id,可以传空值
	_instanceId string
	// 项目所在平台上的gameid
	_gameId string
	// projectKey
	_gameProjectKey string
	// user的外部唯一Id
	_customerAccountId string
	// 消息体
	_message string
	// user的mpaccountId
	_mpAccountId int64
	// user的类型
	_accountType int64
}

// NewAlibabacgamempmpsessionsendmessagetogameRequest 初始化AlibabacgamempmpsessionsendmessagetogameAPIRequest对象
func NewAlibabacgamempmpsessionsendmessagetogameRequest() *AlibabacgamempmpsessionsendmessagetogameAPIRequest {
	return &AlibabacgamempmpsessionsendmessagetogameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.mp.mpsession.sendmessagetogame"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备Id， 可以传空值
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetUserId is UserId Setter
// 调用者的userid，若有userAccessToken,可以传空值
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserToken is UserToken Setter
// 目前没有使用，可以传空值
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetUserToken(_userToken string) error {
	r._userToken = _userToken
	r.Set("user_token", _userToken)
	return nil
}

// GetUserToken UserToken Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetUserToken() string {
	return r._userToken
}

// SetUserAccessToken is UserAccessToken Setter
// 登录后获取的accessToken
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetUserAccessToken(_userAccessToken string) error {
	r._userAccessToken = _userAccessToken
	r.Set("user_access_token", _userAccessToken)
	return nil
}

// GetUserAccessToken UserAccessToken Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetUserAccessToken() string {
	return r._userAccessToken
}

// SetInstanceId is InstanceId Setter
// ecs上的实例id,可以传空值
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetInstanceId(_instanceId string) error {
	r._instanceId = _instanceId
	r.Set("instance_id", _instanceId)
	return nil
}

// GetInstanceId InstanceId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetInstanceId() string {
	return r._instanceId
}

// SetGameId is GameId Setter
// 项目所在平台上的gameid
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetGameId(_gameId string) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetGameId() string {
	return r._gameId
}

// SetGameProjectKey is GameProjectKey Setter
// projectKey
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetGameProjectKey(_gameProjectKey string) error {
	r._gameProjectKey = _gameProjectKey
	r.Set("game_project_key", _gameProjectKey)
	return nil
}

// GetGameProjectKey GameProjectKey Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetGameProjectKey() string {
	return r._gameProjectKey
}

// SetCustomerAccountId is CustomerAccountId Setter
// user的外部唯一Id
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetCustomerAccountId(_customerAccountId string) error {
	r._customerAccountId = _customerAccountId
	r.Set("customer_account_id", _customerAccountId)
	return nil
}

// GetCustomerAccountId CustomerAccountId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetCustomerAccountId() string {
	return r._customerAccountId
}

// SetMessage is Message Setter
// 消息体
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetMessage(_message string) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetMessage() string {
	return r._message
}

// SetMpAccountId is MpAccountId Setter
// user的mpaccountId
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetMpAccountId(_mpAccountId int64) error {
	r._mpAccountId = _mpAccountId
	r.Set("mp_account_id", _mpAccountId)
	return nil
}

// GetMpAccountId MpAccountId Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetMpAccountId() int64 {
	return r._mpAccountId
}

// SetAccountType is AccountType Setter
// user的类型
func (r *AlibabacgamempmpsessionsendmessagetogameAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabacgamempmpsessionsendmessagetogameAPIRequest) GetAccountType() int64 {
	return r._accountType
}
