package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaOrderEffectImportAPIRequest 天下秀订单数据导入 API请求
// alibaba.pictures.dengta.order.effect.import
//
// 提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
type AlibabaPicturesDengtaOrderEffectImportAPIRequest struct {
	model.Params
	// 微博订单id
	_imsOrderId string
	// 微博链接
	_url string
	// 阅读数
	_readsCount int64
	// 转发数
	_repostsCount int64
	// 评论数
	_commentsCount int64
	// 点赞数
	_attitudesCount int64
	// 微博昵称
	_weiboNick string
	// 粉丝数
	_followersCount int64
	// 传播关键节点
	_nodesTop string
	// 关键路径
	_keyPath string
	// 微博来源
	_weiboSource string
	// 类型分布
	_verifiedType string
	// 性别分布
	_gender string
	// 粉丝分布
	_fansCount string
	// 地域分布
	_location string
	// 关系图
	_graph string
	// 关键词云
	_words string
	// 每小时转发量
	_repostNumPerHour string
	// 点赞量每小时趋势图
	_attitudesNumPerHour string
	// 是否成功
	_isSuccess int64
	// 失败原因
	_failReason string
}

// NewAlibabaPicturesDengtaOrderEffectImportRequest 初始化AlibabaPicturesDengtaOrderEffectImportAPIRequest对象
func NewAlibabaPicturesDengtaOrderEffectImportRequest() *AlibabaPicturesDengtaOrderEffectImportAPIRequest {
	return &AlibabaPicturesDengtaOrderEffectImportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.order.effect.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ImsOrderId Setter
// 微博订单id
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// Get ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// Set is Url Setter
// 微博链接
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// Get Url Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetUrl() string {
	return r._url
}

// Set is ReadsCount Setter
// 阅读数
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetReadsCount(_readsCount int64) error {
	r._readsCount = _readsCount
	r.Set("reads_count", _readsCount)
	return nil
}

// Get ReadsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetReadsCount() int64 {
	return r._readsCount
}

// Set is RepostsCount Setter
// 转发数
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetRepostsCount(_repostsCount int64) error {
	r._repostsCount = _repostsCount
	r.Set("reposts_count", _repostsCount)
	return nil
}

// Get RepostsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetRepostsCount() int64 {
	return r._repostsCount
}

// Set is CommentsCount Setter
// 评论数
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetCommentsCount(_commentsCount int64) error {
	r._commentsCount = _commentsCount
	r.Set("comments_count", _commentsCount)
	return nil
}

// Get CommentsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetCommentsCount() int64 {
	return r._commentsCount
}

// Set is AttitudesCount Setter
// 点赞数
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetAttitudesCount(_attitudesCount int64) error {
	r._attitudesCount = _attitudesCount
	r.Set("attitudes_count", _attitudesCount)
	return nil
}

// Get AttitudesCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetAttitudesCount() int64 {
	return r._attitudesCount
}

// Set is WeiboNick Setter
// 微博昵称
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetWeiboNick(_weiboNick string) error {
	r._weiboNick = _weiboNick
	r.Set("weibo_nick", _weiboNick)
	return nil
}

// Get WeiboNick Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetWeiboNick() string {
	return r._weiboNick
}

// Set is FollowersCount Setter
// 粉丝数
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetFollowersCount(_followersCount int64) error {
	r._followersCount = _followersCount
	r.Set("followers_count", _followersCount)
	return nil
}

// Get FollowersCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetFollowersCount() int64 {
	return r._followersCount
}

// Set is NodesTop Setter
// 传播关键节点
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetNodesTop(_nodesTop string) error {
	r._nodesTop = _nodesTop
	r.Set("nodes_top", _nodesTop)
	return nil
}

// Get NodesTop Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetNodesTop() string {
	return r._nodesTop
}

// Set is KeyPath Setter
// 关键路径
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetKeyPath(_keyPath string) error {
	r._keyPath = _keyPath
	r.Set("key_path", _keyPath)
	return nil
}

// Get KeyPath Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetKeyPath() string {
	return r._keyPath
}

// Set is WeiboSource Setter
// 微博来源
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetWeiboSource(_weiboSource string) error {
	r._weiboSource = _weiboSource
	r.Set("weibo_source", _weiboSource)
	return nil
}

// Get WeiboSource Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetWeiboSource() string {
	return r._weiboSource
}

// Set is VerifiedType Setter
// 类型分布
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetVerifiedType(_verifiedType string) error {
	r._verifiedType = _verifiedType
	r.Set("verified_type", _verifiedType)
	return nil
}

// Get VerifiedType Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetVerifiedType() string {
	return r._verifiedType
}

// Set is Gender Setter
// 性别分布
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// Get Gender Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetGender() string {
	return r._gender
}

// Set is FansCount Setter
// 粉丝分布
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetFansCount(_fansCount string) error {
	r._fansCount = _fansCount
	r.Set("fans_count", _fansCount)
	return nil
}

// Get FansCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetFansCount() string {
	return r._fansCount
}

// Set is Location Setter
// 地域分布
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetLocation(_location string) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// Get Location Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetLocation() string {
	return r._location
}

// Set is Graph Setter
// 关系图
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetGraph(_graph string) error {
	r._graph = _graph
	r.Set("graph", _graph)
	return nil
}

// Get Graph Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetGraph() string {
	return r._graph
}

// Set is Words Setter
// 关键词云
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetWords(_words string) error {
	r._words = _words
	r.Set("words", _words)
	return nil
}

// Get Words Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetWords() string {
	return r._words
}

// Set is RepostNumPerHour Setter
// 每小时转发量
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetRepostNumPerHour(_repostNumPerHour string) error {
	r._repostNumPerHour = _repostNumPerHour
	r.Set("repost_num_per_hour", _repostNumPerHour)
	return nil
}

// Get RepostNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetRepostNumPerHour() string {
	return r._repostNumPerHour
}

// Set is AttitudesNumPerHour Setter
// 点赞量每小时趋势图
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetAttitudesNumPerHour(_attitudesNumPerHour string) error {
	r._attitudesNumPerHour = _attitudesNumPerHour
	r.Set("attitudes_num_per_hour", _attitudesNumPerHour)
	return nil
}

// Get AttitudesNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetAttitudesNumPerHour() string {
	return r._attitudesNumPerHour
}

// Set is IsSuccess Setter
// 是否成功
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetIsSuccess(_isSuccess int64) error {
	r._isSuccess = _isSuccess
	r.Set("is_success", _isSuccess)
	return nil
}

// Get IsSuccess Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetIsSuccess() int64 {
	return r._isSuccess
}

// Set is FailReason Setter
// 失败原因
func (r *AlibabaPicturesDengtaOrderEffectImportAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// Get FailReason Getter
func (r AlibabaPicturesDengtaOrderEffectImportAPIRequest) GetFailReason() string {
	return r._failReason
}
