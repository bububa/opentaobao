package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomePictureSync 图片数据同步
// alibaba.alihouse.newhome.picture.sync
//
// 图片数据同步
func AlibabaAlihouseNewhomePictureSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomePictureSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomePictureSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomePictureSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
