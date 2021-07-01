package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductListAPIRequest
商品查询 API请求
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。 */
type AlibabaIcbuProductListAPIRequest struct {
	model.Params
	// 类目ID
	_categoryId int64
	// 商品名称，支持模糊匹配
	_subject string
	// 当前页
	_currentPage int64
	// 每页大小，最大30
	_pageSize int64
	// 商品语种，目前只支持ENGLISH
	_language string
	// 商品三级分组id，可选填。若填写-1 则表示取回的商品没有三级分组，不填入代表取回的商品不关心它的三级分组，填写对应的group id将返回这个分组下的商品
	_groupId3 int64
	// 商品二级分组id，可选填。若填写-1 则表示取回的商品没有二级分组，不填入代表取回的商品不关系它的二级分组，填写对应的group id将返回这个分组下的商品
	_groupId2 int64
	// 商品一级分组id，可选填。若填写0 则表示取回的商品没有一级分组，不填入代表取回的商品不关心它的一级分组，填写对应的group id将返回这个分组下的商品
	_groupId1 int64
	// 商品明文id
	_id int64
	// 最晚修改时间，格式yyyy-MM-dd HH:mm:ss
	_gmtModifiedTo string
	// 最早修改时间，格式yyyy-MM-dd HH:mm:ss
	_gmtModifiedFrom string
}

// New
