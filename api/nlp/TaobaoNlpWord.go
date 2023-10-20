package nlp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlp"
)

// TaobaoNlpWord 文本语言词法分析
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
func TaobaoNlpWord(clt *core.SDKClient, req *nlp.TaobaoNlpWordAPIRequest, resp *nlp.TaobaoNlpWordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
