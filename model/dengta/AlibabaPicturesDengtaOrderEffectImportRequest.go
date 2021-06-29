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
    imsOrderId   string
    // 微博链接
    url   string
    // 阅读数
    readsCount   int64
    // 转发数
    repostsCount   int64
    // 评论数
    commentsCount   int64
    // 点赞数
    attitudesCount   int64
    // 微博昵称
    weiboNick   string
    // 粉丝数
    followersCount   int64
    // 传播关键节点
    nodesTop   string
    // 关键路径
    keyPath   string
    // 微博来源
    weiboSource   string
    // 类型分布
    verifiedType   string
    // 性别分布
    gender   string
    // 粉丝分布
    fansCount   string
    // 地域分布
    location   string
    // 关系图
    graph   string
    // 关键词云
    words   string
    // 每小时转发量
    repostNumPerHour   string
    // 点赞量每小时趋势图
    attitudesNumPerHour   string
    // 是否成功
    isSuccess   int64
    // 失败原因
    failReason   string
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
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetImsOrderId(imsOrderId string) error {
    r.imsOrderId = imsOrderId
    r.Set("ims_order_id", imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetImsOrderId() string {
    return r.imsOrderId
}
// Url Setter
// 微博链接
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetUrl() string {
    return r.url
}
// ReadsCount Setter
// 阅读数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetReadsCount(readsCount int64) error {
    r.readsCount = readsCount
    r.Set("reads_count", readsCount)
    return nil
}

// ReadsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetReadsCount() int64 {
    return r.readsCount
}
// RepostsCount Setter
// 转发数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetRepostsCount(repostsCount int64) error {
    r.repostsCount = repostsCount
    r.Set("reposts_count", repostsCount)
    return nil
}

// RepostsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetRepostsCount() int64 {
    return r.repostsCount
}
// CommentsCount Setter
// 评论数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetCommentsCount(commentsCount int64) error {
    r.commentsCount = commentsCount
    r.Set("comments_count", commentsCount)
    return nil
}

// CommentsCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetCommentsCount() int64 {
    return r.commentsCount
}
// AttitudesCount Setter
// 点赞数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetAttitudesCount(attitudesCount int64) error {
    r.attitudesCount = attitudesCount
    r.Set("attitudes_count", attitudesCount)
    return nil
}

// AttitudesCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetAttitudesCount() int64 {
    return r.attitudesCount
}
// WeiboNick Setter
// 微博昵称
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWeiboNick(weiboNick string) error {
    r.weiboNick = weiboNick
    r.Set("weibo_nick", weiboNick)
    return nil
}

// WeiboNick Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWeiboNick() string {
    return r.weiboNick
}
// FollowersCount Setter
// 粉丝数
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFollowersCount(followersCount int64) error {
    r.followersCount = followersCount
    r.Set("followers_count", followersCount)
    return nil
}

// FollowersCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFollowersCount() int64 {
    return r.followersCount
}
// NodesTop Setter
// 传播关键节点
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetNodesTop(nodesTop string) error {
    r.nodesTop = nodesTop
    r.Set("nodes_top", nodesTop)
    return nil
}

// NodesTop Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetNodesTop() string {
    return r.nodesTop
}
// KeyPath Setter
// 关键路径
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetKeyPath(keyPath string) error {
    r.keyPath = keyPath
    r.Set("key_path", keyPath)
    return nil
}

// KeyPath Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetKeyPath() string {
    return r.keyPath
}
// WeiboSource Setter
// 微博来源
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWeiboSource(weiboSource string) error {
    r.weiboSource = weiboSource
    r.Set("weibo_source", weiboSource)
    return nil
}

// WeiboSource Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWeiboSource() string {
    return r.weiboSource
}
// VerifiedType Setter
// 类型分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetVerifiedType(verifiedType string) error {
    r.verifiedType = verifiedType
    r.Set("verified_type", verifiedType)
    return nil
}

// VerifiedType Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetVerifiedType() string {
    return r.verifiedType
}
// Gender Setter
// 性别分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetGender() string {
    return r.gender
}
// FansCount Setter
// 粉丝分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFansCount(fansCount string) error {
    r.fansCount = fansCount
    r.Set("fans_count", fansCount)
    return nil
}

// FansCount Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFansCount() string {
    return r.fansCount
}
// Location Setter
// 地域分布
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetLocation(location string) error {
    r.location = location
    r.Set("location", location)
    return nil
}

// Location Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetLocation() string {
    return r.location
}
// Graph Setter
// 关系图
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetGraph(graph string) error {
    r.graph = graph
    r.Set("graph", graph)
    return nil
}

// Graph Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetGraph() string {
    return r.graph
}
// Words Setter
// 关键词云
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetWords(words string) error {
    r.words = words
    r.Set("words", words)
    return nil
}

// Words Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetWords() string {
    return r.words
}
// RepostNumPerHour Setter
// 每小时转发量
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetRepostNumPerHour(repostNumPerHour string) error {
    r.repostNumPerHour = repostNumPerHour
    r.Set("repost_num_per_hour", repostNumPerHour)
    return nil
}

// RepostNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetRepostNumPerHour() string {
    return r.repostNumPerHour
}
// AttitudesNumPerHour Setter
// 点赞量每小时趋势图
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetAttitudesNumPerHour(attitudesNumPerHour string) error {
    r.attitudesNumPerHour = attitudesNumPerHour
    r.Set("attitudes_num_per_hour", attitudesNumPerHour)
    return nil
}

// AttitudesNumPerHour Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetAttitudesNumPerHour() string {
    return r.attitudesNumPerHour
}
// IsSuccess Setter
// 是否成功
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetIsSuccess(isSuccess int64) error {
    r.isSuccess = isSuccess
    r.Set("is_success", isSuccess)
    return nil
}

// IsSuccess Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetIsSuccess() int64 {
    return r.isSuccess
}
// FailReason Setter
// 失败原因
func (r *AlibabaPicturesDengtaOrderEffectImportRequest) SetFailReason(failReason string) error {
    r.failReason = failReason
    r.Set("fail_reason", failReason)
    return nil
}

// FailReason Getter
func (r AlibabaPicturesDengtaOrderEffectImportRequest) GetFailReason() string {
    return r.failReason
}
