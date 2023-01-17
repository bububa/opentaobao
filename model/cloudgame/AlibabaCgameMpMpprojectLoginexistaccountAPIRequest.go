package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameMpMpprojectLoginexistaccountAPIRequest 登录存在账号 API请求
// alibaba.cgame.mp.mpproject.loginexistaccount
//
// 发送消息给游戏
type AlibabaCgameMpMpprojectLoginexistaccountAPIRequest struct {
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

// NewAlibabaCgameMpMpprojectLoginexistaccountRequest 初始化AlibabaCgameMpMpprojectLoginexistaccountAPIRequest对象
func NewAlibabaCgameMpMpprojectLoginexistaccountRequest() *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest {
	return &AlibabaCgameMpMpprojectLoginexistaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.mp.mpproject.loginexistaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备Id， 可以传空值
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetUserId is UserId Setter
// 调用者的userid，若有userAccessToken,可以传空值
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserToken is UserToken Setter
// 目前没有使用，可以传空值
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetUserToken(_userToken string) error {
	r._userToken = _userToken
	r.Set("user_token", _userToken)
	return nil
}

// GetUserToken UserToken Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetUserToken() string {
	return r._userToken
}

// SetUserAccessToken is UserAccessToken Setter
// 登录后获取的accessToken
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetUserAccessToken(_userAccessToken string) error {
	r._userAccessToken = _userAccessToken
	r.Set("user_access_token", _userAccessToken)
	return nil
}

// GetUserAccessToken UserAccessToken Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetUserAccessToken() string {
	return r._userAccessToken
}

// SetInstanceId is InstanceId Setter
// ecs上的实例id,可以传空值
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetInstanceId(_instanceId string) error {
	r._instanceId = _instanceId
	r.Set("instance_id", _instanceId)
	return nil
}

// GetInstanceId InstanceId Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetInstanceId() string {
	return r._instanceId
}

// SetGameId is GameId Setter
// 项目所在平台上的gameid
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetGameId(_gameId string) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetGameId() string {
	return r._gameId
}

// SetGameProjectKey is GameProjectKey Setter
// projectKey
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetGameProjectKey(_gameProjectKey string) error {
	r._gameProjectKey = _gameProjectKey
	r.Set("game_project_key", _gameProjectKey)
	return nil
}

// GetGameProjectKey GameProjectKey Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetGameProjectKey() string {
	return r._gameProjectKey
}

// SetCustomerAccountId is CustomerAccountId Setter
// user的外部唯一Id
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetCustomerAccountId(_customerAccountId string) error {
	r._customerAccountId = _customerAccountId
	r.Set("customer_account_id", _customerAccountId)
	return nil
}

// GetCustomerAccountId CustomerAccountId Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetCustomerAccountId() string {
	return r._customerAccountId
}

// SetAccountType is AccountType Setter
// user的类型
func (r *AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabaCgameMpMpprojectLoginexistaccountAPIRequest) GetAccountType() int64 {
	return r._accountType
}
