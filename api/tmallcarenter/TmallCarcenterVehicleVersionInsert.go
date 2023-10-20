package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// Tmallcarcentervehicleversioninsert 汽车EPC版本压缩库新增接口
// tmall.carcenter.vehicle.version.insert
//
// 汽车EPC版本压缩库新增接口
func Tmallcarcentervehicleversioninsert(clt *core.SDKClient, req *tmallcarenter.TmallcarcentervehicleversioninsertAPIRequest, session string) (*tmallcarenter.TmallcarcentervehicleversioninsertAPIResponse, error) {
	var resp tmallcarenter.TmallcarcentervehicleversioninsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
