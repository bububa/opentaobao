package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKfcKeywordSearchAPIRequest
关键词过滤匹配 API请求
taobao.kfc.keyword.search

对输入的文本信息进行禁忌关键词匹配，返回匹配的结果 */
type TaobaoKfcKeywordSearchAPIRequest struct {
	model.Params
	// 发布信息的淘宝会员名，可以不传
	_nick string
	// 需要过滤的文本信息
	_content string
	// 应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。<br/><br/>这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。<br/><br/><br/>通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。
	_apply string
}

// New
