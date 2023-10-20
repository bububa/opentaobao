package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtaordereffectimportAPIRequest 天下秀订单数据导入 API请求
// alibaba.pictures.dengta.order.effect.import
//
// 提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
type AlibabapicturesdengtaordereffectimportAPIRequest struct {
	model.Params
	// 微博来源
	_weiboSource string
	// 地域分布
	_location string
	// 微博链接
	_url string
	// 类型分布
	_verifiedType string
	// 传播关键节点
	_nodesTop string
	// 关系图
	_graph string
	// 关键词云
	_words string
	// 关键路径
	_keyPath string
	// 性别分布
	_gender string
	// 点赞量每小时趋势图
	_attitudesNumPerHour string
	// 微博订单id
	_imsOrderId string
	// 失败原因
	_failReason string
	// 每小时转发量
	_repostNumPerHour string
	// 微博昵称
	_weiboNick string
	// 粉丝分布
	_fansCount string
	// 转发数
	_repostsCount int64
	// 点赞数
	_attitudesCount int64
	// 评论数
	_commentsCount int64
	// 是否成功
	_isSuccess int64
	// 阅读数
	_readsCount int64
	// 粉丝数
	_followersCount int64
}

// NewAlibabapicturesdengtaordereffectimportRequest 初始化AlibabapicturesdengtaordereffectimportAPIRequest对象
func NewAlibabapicturesdengtaordereffectimportRequest() *AlibabapicturesdengtaordereffectimportAPIRequest {
	return &AlibabapicturesdengtaordereffectimportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.order.effect.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWeiboSource is WeiboSource Setter
// 微博来源
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetWeiboSource(_weiboSource string) error {
	r._weiboSource = _weiboSource
	r.Set("weibo_source", _weiboSource)
	return nil
}

// GetWeiboSource WeiboSource Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetWeiboSource() string {
	return r._weiboSource
}

// SetLocation is Location Setter
// 地域分布
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetLocation(_location string) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetLocation() string {
	return r._location
}

// SetUrl is Url Setter
// 微博链接
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetUrl() string {
	return r._url
}

// SetVerifiedType is VerifiedType Setter
// 类型分布
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetVerifiedType(_verifiedType string) error {
	r._verifiedType = _verifiedType
	r.Set("verified_type", _verifiedType)
	return nil
}

// GetVerifiedType VerifiedType Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetVerifiedType() string {
	return r._verifiedType
}

// SetNodesTop is NodesTop Setter
// 传播关键节点
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetNodesTop(_nodesTop string) error {
	r._nodesTop = _nodesTop
	r.Set("nodes_top", _nodesTop)
	return nil
}

// GetNodesTop NodesTop Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetNodesTop() string {
	return r._nodesTop
}

// SetGraph is Graph Setter
// 关系图
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetGraph(_graph string) error {
	r._graph = _graph
	r.Set("graph", _graph)
	return nil
}

// GetGraph Graph Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetGraph() string {
	return r._graph
}

// SetWords is Words Setter
// 关键词云
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetWords(_words string) error {
	r._words = _words
	r.Set("words", _words)
	return nil
}

// GetWords Words Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetWords() string {
	return r._words
}

// SetKeyPath is KeyPath Setter
// 关键路径
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetKeyPath(_keyPath string) error {
	r._keyPath = _keyPath
	r.Set("key_path", _keyPath)
	return nil
}

// GetKeyPath KeyPath Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetKeyPath() string {
	return r._keyPath
}

// SetGender is Gender Setter
// 性别分布
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetGender() string {
	return r._gender
}

// SetAttitudesNumPerHour is AttitudesNumPerHour Setter
// 点赞量每小时趋势图
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetAttitudesNumPerHour(_attitudesNumPerHour string) error {
	r._attitudesNumPerHour = _attitudesNumPerHour
	r.Set("attitudes_num_per_hour", _attitudesNumPerHour)
	return nil
}

// GetAttitudesNumPerHour AttitudesNumPerHour Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetAttitudesNumPerHour() string {
	return r._attitudesNumPerHour
}

// SetImsOrderId is ImsOrderId Setter
// 微博订单id
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// GetImsOrderId ImsOrderId Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// SetFailReason is FailReason Setter
// 失败原因
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetFailReason(_failReason string) error {
	r._failReason = _failReason
	r.Set("fail_reason", _failReason)
	return nil
}

// GetFailReason FailReason Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetFailReason() string {
	return r._failReason
}

// SetRepostNumPerHour is RepostNumPerHour Setter
// 每小时转发量
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetRepostNumPerHour(_repostNumPerHour string) error {
	r._repostNumPerHour = _repostNumPerHour
	r.Set("repost_num_per_hour", _repostNumPerHour)
	return nil
}

// GetRepostNumPerHour RepostNumPerHour Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetRepostNumPerHour() string {
	return r._repostNumPerHour
}

// SetWeiboNick is WeiboNick Setter
// 微博昵称
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetWeiboNick(_weiboNick string) error {
	r._weiboNick = _weiboNick
	r.Set("weibo_nick", _weiboNick)
	return nil
}

// GetWeiboNick WeiboNick Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetWeiboNick() string {
	return r._weiboNick
}

// SetFansCount is FansCount Setter
// 粉丝分布
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetFansCount(_fansCount string) error {
	r._fansCount = _fansCount
	r.Set("fans_count", _fansCount)
	return nil
}

// GetFansCount FansCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetFansCount() string {
	return r._fansCount
}

// SetRepostsCount is RepostsCount Setter
// 转发数
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetRepostsCount(_repostsCount int64) error {
	r._repostsCount = _repostsCount
	r.Set("reposts_count", _repostsCount)
	return nil
}

// GetRepostsCount RepostsCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetRepostsCount() int64 {
	return r._repostsCount
}

// SetAttitudesCount is AttitudesCount Setter
// 点赞数
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetAttitudesCount(_attitudesCount int64) error {
	r._attitudesCount = _attitudesCount
	r.Set("attitudes_count", _attitudesCount)
	return nil
}

// GetAttitudesCount AttitudesCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetAttitudesCount() int64 {
	return r._attitudesCount
}

// SetCommentsCount is CommentsCount Setter
// 评论数
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetCommentsCount(_commentsCount int64) error {
	r._commentsCount = _commentsCount
	r.Set("comments_count", _commentsCount)
	return nil
}

// GetCommentsCount CommentsCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetCommentsCount() int64 {
	return r._commentsCount
}

// SetIsSuccess is IsSuccess Setter
// 是否成功
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetIsSuccess(_isSuccess int64) error {
	r._isSuccess = _isSuccess
	r.Set("is_success", _isSuccess)
	return nil
}

// GetIsSuccess IsSuccess Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetIsSuccess() int64 {
	return r._isSuccess
}

// SetReadsCount is ReadsCount Setter
// 阅读数
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetReadsCount(_readsCount int64) error {
	r._readsCount = _readsCount
	r.Set("reads_count", _readsCount)
	return nil
}

// GetReadsCount ReadsCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetReadsCount() int64 {
	return r._readsCount
}

// SetFollowersCount is FollowersCount Setter
// 粉丝数
func (r *AlibabapicturesdengtaordereffectimportAPIRequest) SetFollowersCount(_followersCount int64) error {
	r._followersCount = _followersCount
	r.Set("followers_count", _followersCount)
	return nil
}

// GetFollowersCount FollowersCount Getter
func (r AlibabapicturesdengtaordereffectimportAPIRequest) GetFollowersCount() int64 {
	return r._followersCount
}
