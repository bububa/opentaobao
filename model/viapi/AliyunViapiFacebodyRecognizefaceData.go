package viapi

// AliyunViapiFacebodyRecognizefaceData 
type AliyunViapiFacebodyRecognizefaceData struct {

    // 返回人脸矩形框，分别是[left, top, width, height], 如有多个人脸，则依次顺延，返回矩形框。如有两个人脸则返回[left1, top1, width1, height1, left2, top2, width2, height2]
    
    FaceRectangles   []int64 `json:"face_rectangles,omitempty" xml:"face_rectangles>int64,omitempty"`
    

    // 人脸识别返回特征维度，目前固定为1024
    
    DenseFeatureLength   int64 `json:"dense_feature_length,omitempty" xml:"dense_feature_length,omitempty"`
    

    // 返回人脸姿态[yaw, pitch, roll]， yaw为左右角度，取值[-90, 90]，pitch为上下角度，取值[-90, 90]， roll为平面旋转角度，取值[-180, 180]，如有多个人脸，则依次顺延
    
    PoseList   []string `json:"pose_list,omitempty" xml:"pose_list>string,omitempty"`
    

    // 左右两个瞳孔的中心点坐标和半径，每个人脸6个浮点数，顺序是[left_iris_cenpt.x, left_iris_cenpt.y, left_iris_radius, right_iris_cenpt.x, right_iris_cenpt.y, right_iris_radis]
    
    Pupils   []string `json:"pupils,omitempty" xml:"pupils>string,omitempty"`
    

    // 0： 女性，1： 男性，如有多个人脸，则依次返回性别
    
    GenderList   []int64 `json:"gender_list,omitempty" xml:"gender_list>int64,omitempty"`
    

    // 人脸识别返回特征；如有多个人脸，则依次顺延，返回特征
    
    DenseFeatures   []string `json:"dense_features,omitempty" xml:"dense_features>string,omitempty"`
    

    // 返回人脸概率, 0-1之间，如有多个人脸，则依次顺延。如有两个人脸则返回[face_prob1, face_prob2]
    
    FaceProbabilityList   []int64 `json:"face_probability_list,omitempty" xml:"face_probability_list>int64,omitempty"`
    

    // 特征点数目，目前固定为105点(顺序：眉毛24点，眼睛32点，鼻子6点，嘴巴34点，外轮廓9点)
    
    LandmarkCount   int64 `json:"landmark_count,omitempty" xml:"landmark_count,omitempty"`
    

    // 年龄0-100，如有多个人脸，依次返回年龄
    
    AgeList   []int64 `json:"age_list,omitempty" xml:"age_list>int64,omitempty"`
    

    // 是否佩戴眼镜，0：无眼镜，1：有眼镜
    
    Glasses   []int64 `json:"glasses,omitempty" xml:"glasses>int64,omitempty"`
    

    // 特征点定位结果，每个人脸返回一组特征点位置，表示方式为（x0, y0, x1, y1, ……）；如有多个人脸，则依次顺延，返回定位浮点数
    
    Landmarks   []string `json:"landmarks,omitempty" xml:"landmarks>string,omitempty"`
    

    // 2种表情，0：中性，1：微笑
    
    Expressions   []int64 `json:"expressions,omitempty" xml:"expressions>int64,omitempty"`
    

    // 检测出来的人脸个数
    
    FaceCount   int64 `json:"face_count,omitempty" xml:"face_count,omitempty"`
    

}
