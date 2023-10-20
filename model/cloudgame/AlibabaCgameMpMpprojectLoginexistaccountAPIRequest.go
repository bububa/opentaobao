package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamempmpprojectloginexistaccountAPIRequest 登录存在账号 API请求
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
type AlibabacgamempmpprojectloginexistaccountAPIRequest struct {
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
	// user的类型
	_accountType int64
}

// NewAlibabacgamempmpprojectloginexistaccountRequest 初始化AlibabacgamempmpprojectloginexistaccountAPIRequest对象
func NewAlibabacgamempmpprojectloginexistaccountRequest() *AlibabacgamempmpprojectloginexistaccountAPIRequest {
	return &AlibabacgamempmpprojectloginexistaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.mp.mpproject.loginexistaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备Id， 可以传空值
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetUserId is UserId Setter
// 调用者的userid，若有userAccessToken,可以传空值
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserToken is UserToken Setter
// 目前没有使用，可以传空值
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetUserToken(_userToken string) error {
	r._userToken = _userToken
	r.Set("user_token", _userToken)
	return nil
}

// GetUserToken UserToken Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetUserToken() string {
	return r._userToken
}

// SetUserAccessToken is UserAccessToken Setter
// 登录后获取的accessToken
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetUserAccessToken(_userAccessToken string) error {
	r._userAccessToken = _userAccessToken
	r.Set("user_access_token", _userAccessToken)
	return nil
}

// GetUserAccessToken UserAccessToken Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetUserAccessToken() string {
	return r._userAccessToken
}

// SetInstanceId is InstanceId Setter
// ecs上的实例id,可以传空值
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetInstanceId(_instanceId string) error {
	r._instanceId = _instanceId
	r.Set("instance_id", _instanceId)
	return nil
}

// GetInstanceId InstanceId Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetInstanceId() string {
	return r._instanceId
}

// SetGameId is GameId Setter
// 项目所在平台上的gameid
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetGameId(_gameId string) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetGameId() string {
	return r._gameId
}

// SetGameProjectKey is GameProjectKey Setter
// projectKey
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetGameProjectKey(_gameProjectKey string) error {
	r._gameProjectKey = _gameProjectKey
	r.Set("game_project_key", _gameProjectKey)
	return nil
}

// GetGameProjectKey GameProjectKey Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetGameProjectKey() string {
	return r._gameProjectKey
}

// SetCustomerAccountId is CustomerAccountId Setter
// user的外部唯一Id
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetCustomerAccountId(_customerAccountId string) error {
	r._customerAccountId = _customerAccountId
	r.Set("customer_account_id", _customerAccountId)
	return nil
}

// GetCustomerAccountId CustomerAccountId Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetCustomerAccountId() string {
	return r._customerAccountId
}

// SetAccountType is AccountType Setter
// user的类型
func (r *AlibabacgamempmpprojectloginexistaccountAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabacgamempmpprojectloginexistaccountAPIRequest) GetAccountType() int64 {
	return r._accountType
}
