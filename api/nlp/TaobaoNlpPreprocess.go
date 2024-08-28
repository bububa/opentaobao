package nlp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlp"
)

// TaobaoNlpPreprocess 文本语言预处理
// taobao.nlp.preprocess
//
// 实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
func TaobaoNlpPreprocess(ctx context.Context, clt *core.SDKClient, req *nlp.TaobaoNlpPreprocessAPIRequest, resp *nlp.TaobaoNlpPreprocessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
