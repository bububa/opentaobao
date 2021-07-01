package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

/* AlibabaItEslEslimageShowimagecommon
对混合云提供的刷图接口
alibaba.it.esl.eslimage.showimagecommon

混合云使用，提供给isv和我们混合云环境部署的应用刷图 */
func AlibabaItEslEslimageShowimagecommon(clt *core.SDKClient, req *ioti.AlibabaItEslEslimageShowimagecommonAPIRequest, session string) (*ioti.AlibabaItEslEslimageShowimagecommonAPIResponse, error) {
	var resp ioti.AlibabaItEslEslimageShowimagecommonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
