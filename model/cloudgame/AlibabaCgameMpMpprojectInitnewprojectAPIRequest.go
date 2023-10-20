package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgamempmpprojectinitnewprojectAPIRequest 创建新的mpproject API请求
// alibaba.cgame.mp.mpproject.initnewproject
//
// 发送消息给游戏
type AlibabacgamempmpprojectinitnewprojectAPIRequest struct {
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
	// customer unique id
	_customerUniqueId string
	// customer platform env
	_customerEnv string
	// customer project id
	_customerProjectId string
	// 是否检查userToken，目前都是0
	_checkUserToken int64
	// 是否只有一个block，目前都是1
	_onlyOneBlock int64
	// 缺省的blockid,默认是0
	_defaultMpBlockId int64
	// mpprojectid,默认是0
	_mpProjectId int64
}

// NewAlibabacgamempmpprojectinitnewprojectRequest 初始化AlibabacgamempmpprojectinitnewprojectAPIRequest对象
func NewAlibabacgamempmpprojectinitnewprojectRequest() *AlibabacgamempmpprojectinitnewprojectAPIRequest {
	return &AlibabacgamempmpprojectinitnewprojectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.mp.mpproject.initnewproject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备Id， 可以传空值
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetUserId is UserId Setter
// 调用者的userid，若有userAccessToken,可以传空值
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetUserId() string {
	return r._userId
}

// SetUserToken is UserToken Setter
// 目前没有使用，可以传空值
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetUserToken(_userToken string) error {
	r._userToken = _userToken
	r.Set("user_token", _userToken)
	return nil
}

// GetUserToken UserToken Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetUserToken() string {
	return r._userToken
}

// SetUserAccessToken is UserAccessToken Setter
// 登录后获取的accessToken
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetUserAccessToken(_userAccessToken string) error {
	r._userAccessToken = _userAccessToken
	r.Set("user_access_token", _userAccessToken)
	return nil
}

// GetUserAccessToken UserAccessToken Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetUserAccessToken() string {
	return r._userAccessToken
}

// SetInstanceId is InstanceId Setter
// ecs上的实例id,可以传空值
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetInstanceId(_instanceId string) error {
	r._instanceId = _instanceId
	r.Set("instance_id", _instanceId)
	return nil
}

// GetInstanceId InstanceId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetInstanceId() string {
	return r._instanceId
}

// SetGameId is GameId Setter
// 项目所在平台上的gameid
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetGameId(_gameId string) error {
	r._gameId = _gameId
	r.Set("game_id", _gameId)
	return nil
}

// GetGameId GameId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetGameId() string {
	return r._gameId
}

// SetGameProjectKey is GameProjectKey Setter
// projectKey
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetGameProjectKey(_gameProjectKey string) error {
	r._gameProjectKey = _gameProjectKey
	r.Set("game_project_key", _gameProjectKey)
	return nil
}

// GetGameProjectKey GameProjectKey Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetGameProjectKey() string {
	return r._gameProjectKey
}

// SetCustomerUniqueId is CustomerUniqueId Setter
// customer unique id
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetCustomerUniqueId(_customerUniqueId string) error {
	r._customerUniqueId = _customerUniqueId
	r.Set("customer_unique_id", _customerUniqueId)
	return nil
}

// GetCustomerUniqueId CustomerUniqueId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetCustomerUniqueId() string {
	return r._customerUniqueId
}

// SetCustomerEnv is CustomerEnv Setter
// customer platform env
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetCustomerEnv(_customerEnv string) error {
	r._customerEnv = _customerEnv
	r.Set("customer_env", _customerEnv)
	return nil
}

// GetCustomerEnv CustomerEnv Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetCustomerEnv() string {
	return r._customerEnv
}

// SetCustomerProjectId is CustomerProjectId Setter
// customer project id
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetCustomerProjectId(_customerProjectId string) error {
	r._customerProjectId = _customerProjectId
	r.Set("customer_project_id", _customerProjectId)
	return nil
}

// GetCustomerProjectId CustomerProjectId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetCustomerProjectId() string {
	return r._customerProjectId
}

// SetCheckUserToken is CheckUserToken Setter
// 是否检查userToken，目前都是0
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetCheckUserToken(_checkUserToken int64) error {
	r._checkUserToken = _checkUserToken
	r.Set("check_user_token", _checkUserToken)
	return nil
}

// GetCheckUserToken CheckUserToken Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetCheckUserToken() int64 {
	return r._checkUserToken
}

// SetOnlyOneBlock is OnlyOneBlock Setter
// 是否只有一个block，目前都是1
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetOnlyOneBlock(_onlyOneBlock int64) error {
	r._onlyOneBlock = _onlyOneBlock
	r.Set("only_one_block", _onlyOneBlock)
	return nil
}

// GetOnlyOneBlock OnlyOneBlock Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetOnlyOneBlock() int64 {
	return r._onlyOneBlock
}

// SetDefaultMpBlockId is DefaultMpBlockId Setter
// 缺省的blockid,默认是0
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetDefaultMpBlockId(_defaultMpBlockId int64) error {
	r._defaultMpBlockId = _defaultMpBlockId
	r.Set("default_mp_block_id", _defaultMpBlockId)
	return nil
}

// GetDefaultMpBlockId DefaultMpBlockId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetDefaultMpBlockId() int64 {
	return r._defaultMpBlockId
}

// SetMpProjectId is MpProjectId Setter
// mpprojectid,默认是0
func (r *AlibabacgamempmpprojectinitnewprojectAPIRequest) SetMpProjectId(_mpProjectId int64) error {
	r._mpProjectId = _mpProjectId
	r.Set("mp_project_id", _mpProjectId)
	return nil
}

// GetMpProjectId MpProjectId Getter
func (r AlibabacgamempmpprojectinitnewprojectAPIRequest) GetMpProjectId() int64 {
	return r._mpProjectId
}
