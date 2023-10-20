package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautometareceive 汽车说明书元数据上传
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
func Tmallaliautometareceive(clt *core.SDKClient, req *tmallcar.TmallaliautometareceiveAPIRequest, session string) (*tmallcar.TmallaliautometareceiveAPIResponse, error) {
	var resp tmallcar.TmallaliautometareceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
