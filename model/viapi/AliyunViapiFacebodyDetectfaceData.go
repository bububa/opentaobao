package viapi

import (
	"sync"
)

// AliyunViapiFacebodyDetectfaceData 结构体
type AliyunViapiFacebodyDetectfaceData struct {
	// 返回人脸概率, 0-1之间，如有多个人脸，则依次顺延。如有两个人脸则返回[face_prob1, face_prob2]
	FaceProbabilityList []int64 `json:"face_probability_list,omitempty" xml:"face_probability_list>int64,omitempty"`
	// 返回人脸矩形框，分别是[left, top, width, height], 如有多个人脸，则依次顺延，返回矩形框。如有两个人脸则返回[left1, top1, width1, height1, left2, top2, width2, height2]
	FaceRectangles []int64 `json:"face_rectangles,omitempty" xml:"face_rectangles>int64,omitempty"`
	// 特征点定位结果，每个人脸返回一组特征点位置，表示方式为（x0, y0, x1, y1, ……）；如有多个人脸，则依次顺延，返回定位浮点数
	Landmarks []string `json:"landmarks,omitempty" xml:"landmarks>string,omitempty"`
	// 返回人脸姿态[yaw, pitch, roll]， yaw为左右角度，取值[-90, 90]，pitch为上下角度，取值[-90, 90]， roll为平面旋转角度，取值[-180, 180]，如有多个人脸，则依次顺延
	PoseList []string `json:"pose_list,omitempty" xml:"pose_list>string,omitempty"`
	// 左右两个瞳孔的中心点坐标和半径，每个人脸6个浮点数，顺序是[left_iris_cenpt.x, left_iris_cenpt.y, left_iris_radius, right_iris_cenpt.x, right_iris_cenpt.y, right_iris_radis]
	Pupils []string `json:"pupils,omitempty" xml:"pupils>string,omitempty"`
	// 特征点数目，目前固定为105点(顺序：眉毛24点，眼睛32点，鼻子6点，嘴巴34点，外轮廓9点)
	LandmarkCount int64 `json:"landmark_count,omitempty" xml:"landmark_count,omitempty"`
	// 检测出来的人脸个数
	FaceCount int64 `json:"face_count,omitempty" xml:"face_count,omitempty"`
}

var poolAliyunViapiFacebodyDetectfaceData = sync.Pool{
	New: func() any {
		return new(AliyunViapiFacebodyDetectfaceData)
	},
}

// GetAliyunViapiFacebodyDetectfaceData() 从对象池中获取AliyunViapiFacebodyDetectfaceData
func GetAliyunViapiFacebodyDetectfaceData() *AliyunViapiFacebodyDetectfaceData {
	return poolAliyunViapiFacebodyDetectfaceData.Get().(*AliyunViapiFacebodyDetectfaceData)
}

// ReleaseAliyunViapiFacebodyDetectfaceData 释放AliyunViapiFacebodyDetectfaceData
func ReleaseAliyunViapiFacebodyDetectfaceData(v *AliyunViapiFacebodyDetectfaceData) {
	v.FaceProbabilityList = v.FaceProbabilityList[:0]
	v.FaceRectangles = v.FaceRectangles[:0]
	v.Landmarks = v.Landmarks[:0]
	v.PoseList = v.PoseList[:0]
	v.Pupils = v.Pupils[:0]
	v.LandmarkCount = 0
	v.FaceCount = 0
	poolAliyunViapiFacebodyDetectfaceData.Put(v)
}
