package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀订单数据导入 API请求
alibaba.pictures.dengta.order.effect.import

提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
*/
type AlibabaPicturesDengtaOrderEffectImportRequest struct {
    model.Params
    // 微博订单id
    _imsOrderId   string
    // 微博链接
    _url   string
    // 阅读数
    _readsCount   int64
    // 转发数
    _repostsCount   int64
    // 评论数
    _commentsCount   int64
    // 点赞数
    _attitudesCount   int64
    // 微博昵称
    _weiboNick   string
    // 粉丝数
    _followersCount   int64
    // 传播关键节点
    _nodesTop   string
    // 关键路径
    _keyPath   string
    // 微博来源
    _weiboSource   string
    // 类型分布
    _verifiedType   string
    // 性别分布
    _gender   string
    // 粉丝分布
    _fansCount   string
    // 地域分布
    _location   string
    // 关系图
    _graph   string
    // 关键词云
    _words   string
    // 每小时转发量
    _repostNumPerHour   string
    // 点赞量每小时趋势图
    _attitudesNumPerHour   string
    // 是否成功
    _isSuccess   int64
    // 失败原因
    _failReason   string
}

// 初始化AlibabaPicturesDengtaOrderEffectImportRequest对象
func NewAlibabaPicturesDengtaOrderEffectImportRequest() *AlibabaPicturesDengtaOrderEffectImportRequest{
    return &AlibabaPicturesDengtaOrderEffectImportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.order.effect.import"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImsOrderId Setter
// 微博订单id
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetImsOrderId(_imsOrderId string) error {
    r._imsOrderId = _imsOrderId
    r.Set("ims_order_id", _imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetImsOrderId() string {
    return r._imsOrderId
}
// Url Setter
// 微博链接
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetUrl() string {
    return r._url
}
// ReadsCount Setter
// 阅读数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetReadsCount(_readsCount int64) error {
    r._readsCount = _readsCount
    r.Set("reads_count", _readsCount)
    return nil
}

// ReadsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetReadsCount() int64 {
    return r._readsCount
}
// RepostsCount Setter
// 转发数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetRepostsCount(_repostsCount int64) error {
    r._repostsCount = _repostsCount
    r.Set("reposts_count", _repostsCount)
    return nil
}

// RepostsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetRepostsCount() int64 {
    return r._repostsCount
}
// CommentsCount Setter
// 评论数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetCommentsCount(_commentsCount int64) error {
    r._commentsCount = _commentsCount
    r.Set("comments_count", _commentsCount)
    return nil
}

// CommentsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetCommentsCount() int64 {
    return r._commentsCount
}
// AttitudesCount Setter
// 点赞数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetAttitudesCount(_attitudesCount int64) error {
    r._attitudesCount = _attitudesCount
    r.Set("attitudes_count", _attitudesCount)
    return nil
}

// AttitudesCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetAttitudesCount() int64 {
    return r._attitudesCount
}
// WeiboNick Setter
// 微博昵称
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWeiboNick(_weiboNick string) error {
    r._weiboNick = _weiboNick
    r.Set("weibo_nick", _weiboNick)
    return nil
}

// WeiboNick Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWeiboNick() string {
    return r._weiboNick
}
// FollowersCount Setter
// 粉丝数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFollowersCount(_followersCount int64) error {
    r._followersCount = _followersCount
    r.Set("followers_count", _followersCount)
    return nil
}

// FollowersCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFollowersCount() int64 {
    return r._followersCount
}
// NodesTop Setter
// 传播关键节点
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetNodesTop(_nodesTop string) error {
    r._nodesTop = _nodesTop
    r.Set("nodes_top", _nodesTop)
    return nil
}

// NodesTop Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetNodesTop() string {
    return r._nodesTop
}
// KeyPath Setter
// 关键路径
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetKeyPath(_keyPath string) error {
    r._keyPath = _keyPath
    r.Set("key_path", _keyPath)
    return nil
}

// KeyPath Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetKeyPath() string {
    return r._keyPath
}
// WeiboSource Setter
// 微博来源
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWeiboSource(_weiboSource string) error {
    r._weiboSource = _weiboSource
    r.Set("weibo_source", _weiboSource)
    return nil
}

// WeiboSource Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWeiboSource() string {
    return r._weiboSource
}
// VerifiedType Setter
// 类型分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetVerifiedType(_verifiedType string) error {
    r._verifiedType = _verifiedType
    r.Set("verified_type", _verifiedType)
    return nil
}

// VerifiedType Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetVerifiedType() string {
    return r._verifiedType
}
// Gender Setter
// 性别分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetGender(_gender string) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetGender() string {
    return r._gender
}
// FansCount Setter
// 粉丝分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFansCount(_fansCount string) error {
    r._fansCount = _fansCount
    r.Set("fans_count", _fansCount)
    return nil
}

// FansCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFansCount() string {
    return r._fansCount
}
// Location Setter
// 地域分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetLocation(_location string) error {
    r._location = _location
    r.Set("location", _location)
    return nil
}

// Location Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetLocation() string {
    return r._location
}
// Graph Setter
// 关系图
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetGraph(_graph string) error {
    r._graph = _graph
    r.Set("graph", _graph)
    return nil
}

// Graph Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetGraph() string {
    return r._graph
}
// Words Setter
// 关键词云
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWords(_words string) error {
    r._words = _words
    r.Set("words", _words)
    return nil
}

// Words Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWords() string {
    return r._words
}
// RepostNumPerHour Setter
// 每小时转发量
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetRepostNumPerHour(_repostNumPerHour string) error {
    r._repostNumPerHour = _repostNumPerHour
    r.Set("repost_num_per_hour", _repostNumPerHour)
    return nil
}

// RepostNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetRepostNumPerHour() string {
    return r._repostNumPerHour
}
// AttitudesNumPerHour Setter
// 点赞量每小时趋势图
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetAttitudesNumPerHour(_attitudesNumPerHour string) error {
    r._attitudesNumPerHour = _attitudesNumPerHour
    r.Set("attitudes_num_per_hour", _attitudesNumPerHour)
    return nil
}

// AttitudesNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetAttitudesNumPerHour() string {
    return r._attitudesNumPerHour
}
// IsSuccess Setter
// 是否成功
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetIsSuccess(_isSuccess int64) error {
    r._isSuccess = _isSuccess
    r.Set("is_success", _isSuccess)
    return nil
}

// IsSuccess Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetIsSuccess() int64 {
    return r._isSuccess
}
// FailReason Setter
// 失败原因
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFailReason(_failReason string) error {
    r._failReason = _failReason
    r.Set("fail_reason", _failReason)
    return nil
}

// FailReason Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFailReason() string {
    return r._failReason
}
