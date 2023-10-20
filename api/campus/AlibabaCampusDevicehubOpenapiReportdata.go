package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdevicehubopenapireportdata 设备数据上报
// alibaba.campus.devicehub.openapi.reportdata
//
// 设备数据上报
func Alibabacampusdevicehubopenapireportdata(clt *core.SDKClient, req *campus.AlibabacampusdevicehubopenapireportdataAPIRequest, session string) (*campus.AlibabacampusdevicehubopenapireportdataAPIResponse, error) {
	var resp campus.AlibabacampusdevicehubopenapireportdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
