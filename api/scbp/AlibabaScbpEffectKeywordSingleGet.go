package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpeffectkeywordsingleget 单个关键词效果报表
// alibaba.scbp.effect.keyword.single.get
//
// 单个关键词效果报表
func Alibabascbpeffectkeywordsingleget(clt *core.SDKClient, req *scbp.AlibabascbpeffectkeywordsinglegetAPIRequest, session string) (*scbp.AlibabascbpeffectkeywordsinglegetAPIResponse, error) {
	var resp scbp.AlibabascbpeffectkeywordsinglegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
