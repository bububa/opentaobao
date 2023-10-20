package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbainsightrelatedwordsget 获取词的相关词
// taobao.simba.insight.relatedwords.get
//
// 获取给定词的若干相关词，返回结果中越相关的权重越大，排在越前面，根据number参数对返回结果进行截断。
func Taobaosimbainsightrelatedwordsget(clt *core.SDKClient, req *simba.TaobaosimbainsightrelatedwordsgetAPIRequest, session string) (*simba.TaobaosimbainsightrelatedwordsgetAPIResponse, error) {
	var resp simba.TaobaosimbainsightrelatedwordsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
