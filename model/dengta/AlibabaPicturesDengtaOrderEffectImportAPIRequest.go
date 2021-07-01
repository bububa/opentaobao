package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaOrderEffectImportAPIRequest
天下秀订单数据导入 API请求
alibaba.pictures.dengta.order.effect.import

提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统 */
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

// New
