// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20211111

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CFSConfig struct {
	// cfs的实例的ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 存储的路径
	Path *string `json:"Path,omitempty" name:"Path"`
}

type CosPathInfo struct {
	// 存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 路径列表，目前只支持单个
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*string `json:"Paths,omitempty" name:"Paths"`
}

// Predefined struct for user
type CreateDatasetRequestParams struct {
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	AnnotationType *string `json:"AnnotationType,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	AnnotationFormat *string `json:"AnnotationFormat,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitempty" name:"IsSchemaExisted"`
}

type CreateDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	DatasetName *string `json:"DatasetName,omitempty" name:"DatasetName"`

	// 数据集类型:
	// TYPE_DATASET_TEXT，文本
	// TYPE_DATASET_IMAGE，图片
	// TYPE_DATASET_TABLE，表格
	// TYPE_DATASET_OTHER，其他
	DatasetType *string `json:"DatasetType,omitempty" name:"DatasetType"`

	// 数据源cos路径
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitempty" name:"StorageDataPath"`

	// 数据集标签cos存储路径
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitempty" name:"StorageLabelPath"`

	// 数据集标签
	DatasetTags []*Tag `json:"DatasetTags,omitempty" name:"DatasetTags"`

	// 数据集标注状态:
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ANNOTATED，已标注
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 标注类型:
	// ANNOTATION_TYPE_CLASSIFICATION，图片分类
	// ANNOTATION_TYPE_DETECTION，目标检测
	// ANNOTATION_TYPE_SEGMENTATION，图片分割
	// ANNOTATION_TYPE_TRACKING，目标跟踪
	AnnotationType *string `json:"AnnotationType,omitempty" name:"AnnotationType"`

	// 标注格式:
	// ANNOTATION_FORMAT_TI，TI平台格式
	// ANNOTATION_FORMAT_PASCAL，Pascal Voc
	// ANNOTATION_FORMAT_COCO，COCO
	// ANNOTATION_FORMAT_FILE，文件目录结构
	AnnotationFormat *string `json:"AnnotationFormat,omitempty" name:"AnnotationFormat"`

	// 表头信息
	SchemaInfos []*SchemaInfo `json:"SchemaInfos,omitempty" name:"SchemaInfos"`

	// 数据是否存在表头
	IsSchemaExisted *bool `json:"IsSchemaExisted,omitempty" name:"IsSchemaExisted"`
}

func (r *CreateDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetName")
	delete(f, "DatasetType")
	delete(f, "StorageDataPath")
	delete(f, "StorageLabelPath")
	delete(f, "DatasetTags")
	delete(f, "AnnotationStatus")
	delete(f, "AnnotationType")
	delete(f, "AnnotationFormat")
	delete(f, "SchemaInfos")
	delete(f, "IsSchemaExisted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatasetResponseParams struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDatasetResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatasetResponseParams `json:"Response"`
}

func (r *CreateDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelRequestParams struct {
	// 导入方式（MODEL/VERSION）
	ImportMethod *string `json:"ImportMethod,omitempty" name:"ImportMethod"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitempty" name:"TrainingModelCosPath"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS/AUTO_ML）
	TrainingModelSource *string `json:"TrainingModelSource,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID
	AutoMLTaskId *string `json:"AutoMLTaskId,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML）
	ModelFormat *string `json:"ModelFormat,omitempty" name:"ModelFormat"`
}

type CreateTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 导入方式（MODEL/VERSION）
	ImportMethod *string `json:"ImportMethod,omitempty" name:"ImportMethod"`

	// 模型来源cos目录，以/结尾
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitempty" name:"TrainingModelCosPath"`

	// 推理环境来源（SYSTEM/CUSTOM）
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	TrainingModelName *string `json:"TrainingModelName,omitempty" name:"TrainingModelName"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 算法框架 （PYTORCH/TENSORFLOW/DETECTRON2/PMML)
	AlgorithmFramework *string `json:"AlgorithmFramework,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitempty" name:"ReasoningEnvironment"`

	// 训练指标，最多支持1000字符
	TrainingModelIndex *string `json:"TrainingModelIndex,omitempty" name:"TrainingModelIndex"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitempty" name:"TrainingModelVersion"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitempty" name:"ReasoningImageInfo"`

	// 模型移动方式（CUT/COPY）
	ModelMoveMode *string `json:"ModelMoveMode,omitempty" name:"ModelMoveMode"`

	// 训练任务ID
	TrainingJobId *string `json:"TrainingJobId,omitempty" name:"TrainingJobId"`

	// 模型ID（导入新模型不需要，导入新版本需要）
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 模型存储cos目录
	ModelOutputPath *CosPathInfo `json:"ModelOutputPath,omitempty" name:"ModelOutputPath"`

	// 模型来源 （JOB/COS/AUTO_ML）
	TrainingModelSource *string `json:"TrainingModelSource,omitempty" name:"TrainingModelSource"`

	// 模型偏好
	TrainingPreference *string `json:"TrainingPreference,omitempty" name:"TrainingPreference"`

	// 自动学习任务ID
	AutoMLTaskId *string `json:"AutoMLTaskId,omitempty" name:"AutoMLTaskId"`

	// 任务版本
	TrainingJobVersion *string `json:"TrainingJobVersion,omitempty" name:"TrainingJobVersion"`

	// 模型版本类型；
	// 枚举值：NORMAL(通用)  ACCELERATE(加速)
	// 注意:  默认为NORMAL
	ModelVersionType *string `json:"ModelVersionType,omitempty" name:"ModelVersionType"`

	// 模型格式 （PYTORCH/TORCH_SCRIPT/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML）
	ModelFormat *string `json:"ModelFormat,omitempty" name:"ModelFormat"`
}

func (r *CreateTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImportMethod")
	delete(f, "TrainingModelCosPath")
	delete(f, "ReasoningEnvironmentSource")
	delete(f, "TrainingModelName")
	delete(f, "Tags")
	delete(f, "TrainingJobName")
	delete(f, "AlgorithmFramework")
	delete(f, "ReasoningEnvironment")
	delete(f, "TrainingModelIndex")
	delete(f, "TrainingModelVersion")
	delete(f, "ReasoningImageInfo")
	delete(f, "ModelMoveMode")
	delete(f, "TrainingJobId")
	delete(f, "TrainingModelId")
	delete(f, "ModelOutputPath")
	delete(f, "TrainingModelSource")
	delete(f, "TrainingPreference")
	delete(f, "AutoMLTaskId")
	delete(f, "TrainingJobVersion")
	delete(f, "ModelVersionType")
	delete(f, "ModelFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingModelResponseParams struct {
	// 模型ID，TrainingModel ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingModelResponseParams `json:"Response"`
}

func (r *CreateTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskRequestParams struct {
	// 训练任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitempty" name:"Name"`

	// 训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD
	TrainingMode *string `json:"TrainingMode,omitempty" name:"TrainingMode"`

	// 计费模式，eg：PREPAID预付费，即包年包月；POSTPAID_BY_HOUR按小时后付费
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}]
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitempty" name:"ResourceConfigInfos"`

	// COS代码包路径
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitempty" name:"CodePackagePath"`

	// COS训练输出路径
	Output *CosPathInfo `json:"Output,omitempty" name:"Output"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitempty" name:"LogEnable"`

	// 训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、TENSORFLOW、PYTORCH、LIGHT
	FrameworkName *string `json:"FrameworkName,omitempty" name:"FrameworkName"`

	// 训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15-py3.6-cpu、1.9-py3.6-cuda11.1-gpu
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" name:"FrameworkVersion"`

	// 预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 启动命令信息，默认为sh start.sh
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitempty" name:"StartCmdInfo"`

	// 数据来源，eg：DATASET、COS、CFS、HDFS
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 数据配置
	DataConfigs []*DataConfig `json:"DataConfigs,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CLS日志配置
	LogConfig *LogConfig `json:"LogConfig,omitempty" name:"LogConfig"`

	// 调优参数
	TuningParameters *string `json:"TuningParameters,omitempty" name:"TuningParameters"`

	// 备注，最多500个字
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务名称，不超过60个字符，仅支持中英文、数字、下划线"_"、短横"-"，只能以中英文、数字开头
	Name *string `json:"Name,omitempty" name:"Name"`

	// 训练模式，通过DescribeTrainingFrameworks接口查询，eg：PS_WORKER、DDP、MPI、HOROVOD
	TrainingMode *string `json:"TrainingMode,omitempty" name:"TrainingMode"`

	// 计费模式，eg：PREPAID预付费，即包年包月；POSTPAID_BY_HOUR按小时后付费
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 资源配置，需填写对应算力规格ID和节点数量，算力规格ID查询接口为DescribeBillingSpecsPrice，eg：[{"Role":"WORKER", "InstanceType": "TI.S.MEDIUM.POST", "InstanceNum": 1}]
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitempty" name:"ResourceConfigInfos"`

	// COS代码包路径
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitempty" name:"CodePackagePath"`

	// COS训练输出路径
	Output *CosPathInfo `json:"Output,omitempty" name:"Output"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitempty" name:"LogEnable"`

	// 训练框架名称，通过DescribeTrainingFrameworks接口查询，eg：SPARK、TENSORFLOW、PYTORCH、LIGHT
	FrameworkName *string `json:"FrameworkName,omitempty" name:"FrameworkName"`

	// 训练框架版本，通过DescribeTrainingFrameworks接口查询，eg：1.15-py3.6-cpu、1.9-py3.6-cuda11.1-gpu
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" name:"FrameworkVersion"`

	// 预付费专用资源组ID，通过DescribeBillingResourceGroups接口查询
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 标签配置
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 自定义镜像信息
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 启动命令信息，默认为sh start.sh
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitempty" name:"StartCmdInfo"`

	// 数据来源，eg：DATASET、COS、CFS、HDFS
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 数据配置
	DataConfigs []*DataConfig `json:"DataConfigs,omitempty" name:"DataConfigs"`

	// VPC Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CLS日志配置
	LogConfig *LogConfig `json:"LogConfig,omitempty" name:"LogConfig"`

	// 调优参数
	TuningParameters *string `json:"TuningParameters,omitempty" name:"TuningParameters"`

	// 备注，最多500个字
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "TrainingMode")
	delete(f, "ChargeType")
	delete(f, "ResourceConfigInfos")
	delete(f, "CodePackagePath")
	delete(f, "Output")
	delete(f, "LogEnable")
	delete(f, "FrameworkName")
	delete(f, "FrameworkVersion")
	delete(f, "ResourceGroupId")
	delete(f, "Tags")
	delete(f, "ImageInfo")
	delete(f, "StartCmdInfo")
	delete(f, "DataSource")
	delete(f, "DataConfigs")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LogConfig")
	delete(f, "TuningParameters")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrainingTaskResponseParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrainingTaskResponseParams `json:"Response"`
}

func (r *CreateTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomTrainingData struct {
	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*CustomTrainingMetric `json:"Metrics,omitempty" name:"Metrics"`
}

type CustomTrainingMetric struct {
	// X轴数据类型: TIMESTAMP; EPOCH; STEP
	XType *string `json:"XType,omitempty" name:"XType"`

	// 数据点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*CustomTrainingPoint `json:"Points,omitempty" name:"Points"`
}

type CustomTrainingPoint struct {
	// X值
	XValue *float64 `json:"XValue,omitempty" name:"XValue"`

	// Y值
	YValue *float64 `json:"YValue,omitempty" name:"YValue"`
}

type DataConfig struct {
	// 映射路径
	MappingPath *string `json:"MappingPath,omitempty" name:"MappingPath"`

	// DATASET、COS、CFS、HDFS
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 来自数据集的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSetSource *DataSetConfig `json:"DataSetSource,omitempty" name:"DataSetSource"`

	// 来自cos的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSSource *CosPathInfo `json:"COSSource,omitempty" name:"COSSource"`

	// 来自CFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CFSSource *CFSConfig `json:"CFSSource,omitempty" name:"CFSSource"`

	// 来自HDFS的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	HDFSSource *HDFSConfig `json:"HDFSSource,omitempty" name:"HDFSSource"`
}

type DataPoint struct {
	// 指标名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 值
	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DataSetConfig struct {
	// 数据集ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DatasetGroup struct {
	// 数据集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitempty" name:"DatasetName"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitempty" name:"DatasetTags"`

	// 数据集标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *uint64 `json:"Process,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 外部关联TASKType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitempty" name:"ExternalTaskType"`

	// 数据集大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitempty" name:"DatasetSize"`

	// 数据集数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitempty" name:"FileNum"`

	// 数据集源COS路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitempty" name:"StorageDataPath"`

	// 数据集标签存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitempty" name:"StorageLabelPath"`

	// 数据集版本聚合详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersions []*DatasetInfo `json:"DatasetVersions,omitempty" name:"DatasetVersions"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitempty" name:"DatasetScope"`
}

type DatasetInfo struct {
	// 数据集id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 数据集名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetName *string `json:"DatasetName,omitempty" name:"DatasetName"`

	// 数据集创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 数据集版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetVersion *string `json:"DatasetVersion,omitempty" name:"DatasetVersion"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetType *string `json:"DatasetType,omitempty" name:"DatasetType"`

	// 数据集标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetTags []*Tag `json:"DatasetTags,omitempty" name:"DatasetTags"`

	// 数据集对应标注任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskName *string `json:"DatasetAnnotationTaskName,omitempty" name:"DatasetAnnotationTaskName"`

	// 数据集对应标注任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetAnnotationTaskId *string `json:"DatasetAnnotationTaskId,omitempty" name:"DatasetAnnotationTaskId"`

	// 处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *int64 `json:"Process,omitempty" name:"Process"`

	// 数据集状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetStatus *string `json:"DatasetStatus,omitempty" name:"DatasetStatus"`

	// 错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 数据集创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据集更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 外部任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalTaskType *string `json:"ExternalTaskType,omitempty" name:"ExternalTaskType"`

	// 数据集存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetSize *string `json:"DatasetSize,omitempty" name:"DatasetSize"`

	// 数据集数据数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNum *uint64 `json:"FileNum,omitempty" name:"FileNum"`

	// 数据集源cos 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageDataPath *CosPathInfo `json:"StorageDataPath,omitempty" name:"StorageDataPath"`

	// 数据集输出cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLabelPath *CosPathInfo `json:"StorageLabelPath,omitempty" name:"StorageLabelPath"`

	// 数据集标注状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 数据集类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationType *string `json:"AnnotationType,omitempty" name:"AnnotationType"`

	// 数据集标注格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotationFormat *string `json:"AnnotationFormat,omitempty" name:"AnnotationFormat"`

	// 数据集范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetScope *string `json:"DatasetScope,omitempty" name:"DatasetScope"`
}

// Predefined struct for user
type DeleteDatasetRequestParams struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitempty" name:"DeleteLabelEnable"`
}

type DeleteDatasetRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 是否删除cos标签文件
	DeleteLabelEnable *bool `json:"DeleteLabelEnable,omitempty" name:"DeleteLabelEnable"`
}

func (r *DeleteDatasetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "DeleteLabelEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDatasetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDatasetResponseParams struct {
	// 删除的datasetId
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDatasetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDatasetResponseParams `json:"Response"`
}

func (r *DeleteDatasetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDatasetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`
}

type DeleteTrainingModelRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`
}

func (r *DeleteTrainingModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTrainingModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`
}

type DeleteTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`
}

func (r *DeleteTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingModelVersionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DeleteTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DeleteTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrainingTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrainingTaskResponseParams `json:"Response"`
}

func (r *DeleteTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsRequestParams struct {
	// 资源组类型; 枚举值 TRAIN:训练 INFERENCE:推理
	Type *string `json:"Type,omitempty" name:"Type"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)
	//                     ResourceGroupName (资源组名称列表)
	// Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为30;
	// 注意：小于0则默认为20；大于30则默认为30
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitempty" name:"DontShowInstanceSet"`
}

type DescribeBillingResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 资源组类型; 枚举值 TRAIN:训练 INFERENCE:推理
	Type *string `json:"Type,omitempty" name:"Type"`

	// Filter.Name: 枚举值: ResourceGroupId (资源组id列表)
	//                     ResourceGroupName (资源组名称列表)
	// Filter.Values: 长度为1且Filter.Fuzzy=true时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为5，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 偏移量，默认为0；分页查询起始位置，如：Limit为100，第一页Offset为0，第二页OffSet为100....即每页左边为闭区间
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为30;
	// 注意：小于0则默认为20；大于30则默认为30
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 支持模糊查找资源组id和资源组名
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 是否不展示节点列表; 
	// true: 不展示，false 展示；
	// 默认为false
	DontShowInstanceSet *bool `json:"DontShowInstanceSet,omitempty" name:"DontShowInstanceSet"`
}

func (r *DescribeBillingResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "DontShowInstanceSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingResourceGroupsResponseParams struct {
	// 资源组总数； 注意接口是分页拉取的，total是指资源组总数，不是本次返回中ResourceGroupSet数组的大小
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资源组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupSet []*ResourceGroup `json:"ResourceGroupSet,omitempty" name:"ResourceGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillingResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingResourceGroupsResponseParams `json:"Response"`
}

func (r *DescribeBillingResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceRequestParams struct {
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitempty" name:"SpecsParam"`
}

type DescribeBillingSpecsPriceRequest struct {
	*tchttp.BaseRequest
	
	// 询价参数，支持批量询价
	SpecsParam []*SpecUnit `json:"SpecsParam,omitempty" name:"SpecsParam"`
}

func (r *DescribeBillingSpecsPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpecsParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingSpecsPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingSpecsPriceResponseParams struct {
	// 计费项价格，支持批量返回
	SpecsPrice []*SpecPrice `json:"SpecsPrice,omitempty" name:"SpecsPrice"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBillingSpecsPriceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingSpecsPriceResponseParams `json:"Response"`
}

func (r *DescribeBillingSpecsPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingSpecsPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailStructuredRequestParams struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据条数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDatasetDetailStructuredRequest struct {
	*tchttp.BaseRequest
	
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据条数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDatasetDetailStructuredRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailStructuredRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetDetailStructuredRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailStructuredResponseParams struct {
	// 数据总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格头信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnNames []*string `json:"ColumnNames,omitempty" name:"ColumnNames"`

	// 表格内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowItems []*RowItem `json:"RowItems,omitempty" name:"RowItems"`

	// 文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowTexts []*string `json:"RowTexts,omitempty" name:"RowTexts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatasetDetailStructuredResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetDetailStructuredResponseParams `json:"Response"`
}

func (r *DescribeDatasetDetailStructuredResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailStructuredResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailUnstructuredRequestParams struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回个数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤参数，对应标签值
	LabelList []*string `json:"LabelList,omitempty" name:"LabelList"`

	// 标注状态过滤参数:
	// STATUS_ANNOTATED，已标注
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ALL，全部
	// 默认为STATUS_ALL
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 数据集ID列表
	DatasetIds []*string `json:"DatasetIds,omitempty" name:"DatasetIds"`
}

type DescribeDatasetDetailUnstructuredRequest struct {
	*tchttp.BaseRequest
	
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回个数，默认20，目前最大支持2000条数据
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤参数，对应标签值
	LabelList []*string `json:"LabelList,omitempty" name:"LabelList"`

	// 标注状态过滤参数:
	// STATUS_ANNOTATED，已标注
	// STATUS_NON_ANNOTATED，未标注
	// STATUS_ALL，全部
	// 默认为STATUS_ALL
	AnnotationStatus *string `json:"AnnotationStatus,omitempty" name:"AnnotationStatus"`

	// 数据集ID列表
	DatasetIds []*string `json:"DatasetIds,omitempty" name:"DatasetIds"`
}

func (r *DescribeDatasetDetailUnstructuredRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailUnstructuredRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LabelList")
	delete(f, "AnnotationStatus")
	delete(f, "DatasetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetDetailUnstructuredRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetDetailUnstructuredResponseParams struct {
	// 已标注数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnnotatedTotalCount *uint64 `json:"AnnotatedTotalCount,omitempty" name:"AnnotatedTotalCount"`

	// 没有标注数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonAnnotatedTotalCount *uint64 `json:"NonAnnotatedTotalCount,omitempty" name:"NonAnnotatedTotalCount"`

	// 过滤数据总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterTotalCount *uint64 `json:"FilterTotalCount,omitempty" name:"FilterTotalCount"`

	// 过滤数据详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterLabelList []*FilterLabelInfo `json:"FilterLabelList,omitempty" name:"FilterLabelList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatasetDetailUnstructuredResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetDetailUnstructuredResponseParams `json:"Response"`
}

func (r *DescribeDatasetDetailUnstructuredResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetDetailUnstructuredResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsRequestParams struct {
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDatasetsRequest struct {
	*tchttp.BaseRequest
	
	// 数据集id列表
	DatasetIds []*string `json:"DatasetIds,omitempty" name:"DatasetIds"`

	// 数据集查询过滤条件，多个Filter之间的关系为逻辑与（AND）关系，过滤字段Filter.Name，类型为String
	// DatasetName，数据集名称
	// DatasetScope，数据集范围，SCOPE_DATASET_PRIVATE或SCOPE_DATASET_PUBLIC
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤条件
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 排序值，支持Asc或Desc，默认Desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段，支持CreateTime或UpdateTime，默认CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数据个数，默认20，最大支持200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDatasetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasetIds")
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasetsResponseParams struct {
	// 数据集总量（名称维度）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据集按照数据集名称聚合的分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetGroups []*DatasetGroup `json:"DatasetGroups,omitempty" name:"DatasetGroups"`

	// 数据集ID总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasetIdNums *uint64 `json:"DatasetIdNums,omitempty" name:"DatasetIdNums"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatasetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasetsResponseParams `json:"Response"`
}

func (r *DescribeDatasetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesRequestParams struct {

}

type DescribeInferTemplatesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInferTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInferTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInferTemplatesResponseParams struct {
	// 模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkTemplates []*InferTemplateGroup `json:"FrameworkTemplates,omitempty" name:"FrameworkTemplates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInferTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInferTemplatesResponseParams `json:"Response"`
}

func (r *DescribeInferTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInferTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestTrainingMetricsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeLatestTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeLatestTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLatestTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestTrainingMetricsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 最近一次上报的训练指标.每个Metric中只有一个点的数据, 即len(Values) = len(Timestamps) = 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*TrainingMetric `json:"Metrics,omitempty" name:"Metrics"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLatestTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLatestTrainingMetricsResponseParams `json:"Response"`
}

func (r *DescribeLatestTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// 查询哪个服务的事件（可选值为TRAIN, NOTEBOOK, INFER）
	Service *string `json:"Service,omitempty" name:"Service"`

	// 查询哪个Pod的日志（支持结尾通配符*)
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志查询条数，默认值100，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方向（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按哪个字段排序（可选值为Timestamp），默认值为Timestamp
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回
	Context *string `json:"Context,omitempty" name:"Context"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前只支持Key（也就是按关键字过滤日志）
	// 2. Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// 查询哪个服务的事件（可选值为TRAIN, NOTEBOOK, INFER）
	Service *string `json:"Service,omitempty" name:"Service"`

	// 查询哪个Pod的日志（支持结尾通配符*)
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 日志查询开始时间（RFC3339格式的时间字符串），默认值为当前时间的前一个小时
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 日志查询结束时间（RFC3339格式的时间字符串），默认值为当前时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 日志查询条数，默认值100，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序方向（可选值为ASC, DESC ），默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 按哪个字段排序（可选值为Timestamp），默认值为Timestamp
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 日志查询上下文，查询下一页的时候需要回传这个字段，该字段来自本接口的返回
	Context *string `json:"Context,omitempty" name:"Context"`

	// 过滤条件
	// 注意: 
	// 1. Filter.Name：目前只支持Key（也就是按关键字过滤日志）
	// 2. Filter.Values：表示过滤日志的关键字；Values为多个的时候表示同时满足
	// 3. Filter. Negative和Filter. Fuzzy没有使用
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "PodName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	delete(f, "Context")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// 分页的游标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*LogIdentity `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingFrameworksRequestParams struct {

}

type DescribeTrainingFrameworksRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTrainingFrameworksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingFrameworksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingFrameworksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingFrameworksResponseParams struct {
	// 框架信息列表
	FrameworkInfos []*FrameworkInfo `json:"FrameworkInfos,omitempty" name:"FrameworkInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingFrameworksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingFrameworksResponseParams `json:"Response"`
}

func (r *DescribeTrainingFrameworksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingFrameworksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingMetricsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingMetricsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 训练指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CustomTrainingData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingMetricsResponseParams `json:"Response"`
}

func (r *DescribeTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionRequestParams struct {
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`
}

type DescribeTrainingModelVersionRequest struct {
	*tchttp.BaseRequest
	
	// 模型版本ID
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`
}

func (r *DescribeTrainingModelVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionResponseParams struct {
	// 模型版本
	TrainingModelVersion *TrainingModelVersionDTO `json:"TrainingModelVersion,omitempty" name:"TrainingModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsRequestParams struct {
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTrainingModelVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 模型ID
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 过滤条件
	// Filter.Name: 枚举值:
	//     TrainingModelVersionId (模型版本ID)
	//     ModelVersionType (模型版本类型) 其值支持: NORMAL(通用) ACCELERATE (加速)
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTrainingModelVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrainingModelId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelVersionsResponseParams struct {
	// 模型版本列表
	TrainingModelVersions []*TrainingModelVersionDTO `json:"TrainingModelVersions,omitempty" name:"TrainingModelVersions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingModelVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelVersionsResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelsRequestParams struct {
	// 过滤器
	// Filter.Name: 枚举值:
	//     keyword (模型名称)
	//     TrainingModelId (模型ID)
	//     ModelVersionType (模型版本类型) 其值Filter.Values支持: NORMAL(通用) ACCELERATE (加速)
	//     TrainingModelSource (模型来源)  其值Filter.Values支持： JOB/COS/AUTO_ML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	// Filter.Fuzzy取值：true/false，是否支持模糊匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，ASC/DESC，默认DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeTrainingModelsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器
	// Filter.Name: 枚举值:
	//     keyword (模型名称)
	//     TrainingModelId (模型ID)
	//     ModelVersionType (模型版本类型) 其值Filter.Values支持: NORMAL(通用) ACCELERATE (加速)
	//     TrainingModelSource (模型来源)  其值Filter.Values支持： JOB/COS/AUTO_ML
	//     AlgorithmFramework (算法框架) 其值Filter.Values支持：TENSORFLOW/PYTORCH/DETECTRON2
	//     ModelFormat（模型格式）其值Filter.Values支持：
	// TORCH_SCRIPT/PYTORCH/DETECTRON2/SAVED_MODEL/FROZEN_GRAPH/PMML
	// Filter.Values: 当长度为1时，支持模糊查询; 不为1时，精确查询
	// 每次请求的Filters的上限为10，Filter.Values的上限为100
	// Filter.Fuzzy取值：true/false，是否支持模糊匹配
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段，默认CreateTime
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方式，ASC/DESC，默认DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeTrainingModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingModelsResponseParams struct {
	// 模型列表
	TrainingModels []*TrainingModelDTO `json:"TrainingModels,omitempty" name:"TrainingModels"`

	// 模型总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingModelsResponseParams `json:"Response"`
}

func (r *DescribeTrainingModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeTrainingTaskPodsRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTrainingTaskPodsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskPodsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskPodsResponseParams struct {
	// pod名称列表
	PodNames []*string `json:"PodNames,omitempty" name:"PodNames"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskPodsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskPodsResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskPodsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type DescribeTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTaskResponseParams struct {
	// 训练任务详情
	TrainingTaskDetail *TrainingTaskDetail `json:"TrainingTaskDetail,omitempty" name:"TrainingTaskDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTaskResponseParams `json:"Response"`
}

func (r *DescribeTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksRequestParams struct {
	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeTrainingTasksRequest struct {
	*tchttp.BaseRequest
	
	// 过滤器，eg：[{ "Name": "Id", "Values": ["train-23091792777383936"] }]
	// 
	// 取值范围：
	// Name（名称）：task1
	// Id（task ID）：train-23091792777383936
	// Status（状态）：STARTING / RUNNING / STOPPING / STOPPED / FAILED / SUCCEED / SUBMIT_FAILED
	// ChargeType（计费类型）：PREPAID（预付费）/ POSTPAID_BY_HOUR（后付费）
	// CHARGE_STATUS（计费状态）：NOT_BILLING（未开始计费）/ BILLING（计费中）/ ARREARS_STOP（欠费停止）
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤器，eg：[{ "TagKey": "TagKeyA", "TagValue": ["TagValueA"] }]
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为10，最大为50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 输出列表的排列顺序。取值范围：ASC（升序排列）/ DESC（降序排列），默认为DESC
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序的依据字段， 取值范围 "CreateTime" "UpdateTime"
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeTrainingTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrainingTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrainingTasksResponseParams struct {
	// 训练任务集
	TrainingTaskSet []*TrainingTaskSetItem `json:"TrainingTaskSet,omitempty" name:"TrainingTaskSet"`

	// 数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrainingTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrainingTasksResponseParams `json:"Response"`
}

func (r *DescribeTrainingTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrainingTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetectionLabelInfo struct {
	// 点坐标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitempty" name:"Points"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitempty" name:"Labels"`

	// 类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameType *string `json:"FrameType,omitempty" name:"FrameType"`
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段取值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否开启反向查询
	Negative *bool `json:"Negative,omitempty" name:"Negative"`

	// 是否开启模糊匹配
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}

type FilterLabelInfo struct {
	// 数据集id
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`

	// 文件ID
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 文件路径
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 分类标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationLabels []*string `json:"ClassificationLabels,omitempty" name:"ClassificationLabels"`

	// 检测标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetectionLabels []*DetectionLabelInfo `json:"DetectionLabels,omitempty" name:"DetectionLabels"`

	// 分割标签结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentationLabels []*SegmentationInfo `json:"SegmentationLabels,omitempty" name:"SegmentationLabels"`

	// RGB 图片路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RGBPath *string `json:"RGBPath,omitempty" name:"RGBPath"`

	// 标签模板类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelTemplateType *string `json:"LabelTemplateType,omitempty" name:"LabelTemplateType"`

	// 下载url链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 缩略图下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadThumbnailUrl *string `json:"DownloadThumbnailUrl,omitempty" name:"DownloadThumbnailUrl"`

	// 分割结果图片下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadRGBUrl *string `json:"DownloadRGBUrl,omitempty" name:"DownloadRGBUrl"`

	// OCR场景
	// IDENTITY：识别
	// STRUCTURE：智能结构化
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrScene *string `json:"OcrScene,omitempty" name:"OcrScene"`

	// OCR场景标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrLabels []*OcrLabelInfo `json:"OcrLabels,omitempty" name:"OcrLabels"`
}

type FrameworkInfo struct {
	// 框架名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 框架版本以及对应的训练模式
	VersionInfos []*FrameworkVersion `json:"VersionInfos,omitempty" name:"VersionInfos"`
}

type FrameworkVersion struct {
	// 框架版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 训练模式
	TrainingModes []*string `json:"TrainingModes,omitempty" name:"TrainingModes"`
}

type GpuDetail struct {
	// GPU 显卡类型；枚举值: V100 A100 T4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// GPU 显卡数；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type GroupResource struct {
	// CPU核数; 单位为1/1000核，比如100表示0.1核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存；单位为MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 总卡数；GPUDetail 显卡数之和；单位为1/100卡，比如100代表1卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// Gpu详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuDetailSet []*GpuDetail `json:"GpuDetailSet,omitempty" name:"GpuDetailSet"`
}

type HDFSConfig struct {
	// 集群实例ID,实例ID形如: emr-xxxxxxxx
	Id *string `json:"Id,omitempty" name:"Id"`

	// 路径
	Path *string `json:"Path,omitempty" name:"Path"`
}

type ImageInfo struct {
	// 镜像类型：TCR为腾讯云TCR镜像; CCR为腾讯云TCR个人版镜像，PreSet为平台预置镜像
	ImageType *string `json:"ImageType,omitempty" name:"ImageType"`

	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// TCR镜像对应的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryRegion *string `json:"RegistryRegion,omitempty" name:"RegistryRegion"`

	// TCR镜像对应的实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

type InferTemplate struct {
	// 模板ID
	InferTemplateId *string `json:"InferTemplateId,omitempty" name:"InferTemplateId"`

	// 模板镜像
	InferTemplateImage *string `json:"InferTemplateImage,omitempty" name:"InferTemplateImage"`
}

type InferTemplateGroup struct {
	// 算法框架
	// 注意：此字段可能返回 null，表示取不到有效值。
	Framework *string `json:"Framework,omitempty" name:"Framework"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" name:"FrameworkVersion"`

	// 支持的训练框架集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*string `json:"Groups,omitempty" name:"Groups"`

	// 镜像模板参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InferTemplates []*InferTemplate `json:"InferTemplates,omitempty" name:"InferTemplates"`
}

type Instance struct {
	// 资源组节点id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点已用资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *ResourceInfo `json:"UsedResource,omitempty" name:"UsedResource"`

	// 节点总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *ResourceInfo `json:"TotalResource,omitempty" name:"TotalResource"`

	// 节点状态 
	// 注意：此字段为枚举值
	// 说明: 
	// DEPLOYING: 部署中
	// RUNNING: 运行中 
	// DEPLOY_FAILED: 部署失败
	//  RELEASING 释放中 
	// RELEASED：已释放 
	// EXCEPTION：异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 创建人
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 创建时间: 
	// 注意：北京时间，比如: 2021-12-01 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 到期时间
	// 注意：北京时间，比如：2021-12-11 12:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 自动续费标识
	// 注意：此字段为枚举值
	// 说明：
	// NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期通知
	// NOTIFY_AND_AUTO_RENEW：自动续费且到期通知
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：手动续费(取消自动续费)且到期不通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *string `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 计费项ID
	SpecId *string `json:"SpecId,omitempty" name:"SpecId"`

	// 计费项别名
	SpecAlias *string `json:"SpecAlias,omitempty" name:"SpecAlias"`
}

type LogConfig struct {
	// 日志需要投递到cls的日志集
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志需要投递到cls的主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type LogIdentity struct {
	// 单条日志的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 单条日志的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 这条日志对应的Pod名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 日志的时间戳（RFC3339格式的时间字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
}

type MetricData struct {
	// 训练任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 时间戳.unix timestamp,单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 本次上报数据所处的训练周期数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Epoch *int64 `json:"Epoch,omitempty" name:"Epoch"`

	// 本次上报数据所处的训练迭代次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Step *int64 `json:"Step,omitempty" name:"Step"`

	// 训练停止所需的迭代总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSteps *int64 `json:"TotalSteps,omitempty" name:"TotalSteps"`

	// 数据点。数组元素为不同指标的数据。数组长度不超过10。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*DataPoint `json:"Points,omitempty" name:"Points"`
}

type OcrLabelInfo struct {
	// 坐标点围起来的框
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitempty" name:"Points"`

	// 框的形状：
	// FRAME_TYPE_RECTANGLE
	// FRAME_TYPE_POLYGON
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameType *string `json:"FrameType,omitempty" name:"FrameType"`

	// 智能结构化：key区域对应的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 智能结构化：上述key的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 识别：框区域的内容
	// 智能结构化：value区域对应的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 智能结构化：value区域所关联的key 区域的keyID的集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyIdsForValue []*string `json:"KeyIdsForValue,omitempty" name:"KeyIdsForValue"`

	// key或者value区域内容的方向：
	// DIRECTION_VERTICAL
	// DIRECTION_HORIZONTAL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type PointInfo struct {
	// X坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *float64 `json:"X,omitempty" name:"X"`

	// Y坐标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitempty" name:"Y"`
}

// Predefined struct for user
type PushTrainingMetricsRequestParams struct {
	// 指标数据
	Data []*MetricData `json:"Data,omitempty" name:"Data"`
}

type PushTrainingMetricsRequest struct {
	*tchttp.BaseRequest
	
	// 指标数据
	Data []*MetricData `json:"Data,omitempty" name:"Data"`
}

func (r *PushTrainingMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushTrainingMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushTrainingMetricsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PushTrainingMetricsResponse struct {
	*tchttp.BaseResponse
	Response *PushTrainingMetricsResponseParams `json:"Response"`
}

func (r *PushTrainingMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushTrainingMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceConfigInfo struct {
	// 角色，eg：PS、WORKER、DRIVER、EXECUTOR
	Role *string `json:"Role,omitempty" name:"Role"`

	// cpu核数，1000=1核
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，单位为MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// gpu卡类型
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`

	// gpu数
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// 算力规格ID
	// 计算规格 (for后付费)，可选值如下：
	// TI.S.LARGE.POST: 4C8G 
	// TI.S.2XLARGE16.POST:  8C16G 
	// TI.S.2XLARGE32.POST:  8C32G 
	// TI.S.4XLARGE32.POST:  16C32G
	// TI.S.4XLARGE64.POST:  16C64G
	// TI.S.6XLARGE48.POST:  24C48G
	// TI.S.6XLARGE96.POST:  24C96G
	// TI.S.8XLARGE64.POST:  32C64G
	// TI.S.8XLARGE128.POST : 32C128G
	// TI.GN10.2XLARGE40.POST: 8C40G V100*1 
	// TI.GN10.5XLARGE80.POST:  18C80G V100*2 
	// TI.GN10.10XLARGE160.POST :  32C160G V100*4
	// TI.GN10.20XLARGE320.POST :  72C320G V100*8
	// TI.GN7.8XLARGE128.POST: 32C128G T4*1 
	// TI.GN7.10XLARGE160.POST: 40C160G T4*2 
	// TI.GN7.20XLARGE320.POST: 80C32
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 计算节点数
	InstanceNum *uint64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 算力规格名称
	// 计算规格 (for后付费)，可选值如下：
	// 4C8G 
	// 8C16G 
	// 8C32G 
	// 16C32G
	// 6C64G
	// 24C48G
	// 24C96G
	// 32C64G
	// 32C128G
	// 8C40G V100*1 
	// 8C80G V100*2 
	// 32C160G V100*4
	// 72C320G V100*8
	// 32C128G T4*1 
	// 40C160G T4*2 
	// 80C32
	InstanceTypeAlias *string `json:"InstanceTypeAlias,omitempty" name:"InstanceTypeAlias"`
}

type ResourceGroup struct {
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`

	// 可用节点个数(运行中的节点)
	FreeInstance *uint64 `json:"FreeInstance,omitempty" name:"FreeInstance"`

	// 总节点个数(所有节点)
	TotalInstance *uint64 `json:"TotalInstance,omitempty" name:"TotalInstance"`

	// 资资源组已用的资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedResource *GroupResource `json:"UsedResource,omitempty" name:"UsedResource"`

	// 资源组总资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalResource *GroupResource `json:"TotalResource,omitempty" name:"TotalResource"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitempty" name:"TagSet"`
}

type ResourceInfo struct {
	// 处理器资源, 单位为1/1000核
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存资源, 单位为1M
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// Gpu卡个数资源, 单位为0.01单位的GpuType.
	// Gpu=100表示使用了“一张”gpu卡, 但此处的“一张”卡有可能是虚拟化后的1/4卡, 也有可能是整张卡. 取决于实例的机型
	// 例1 实例的机型带有1张虚拟gpu卡, 每张虚拟gpu卡对应1/4张实际T4卡, 则此时 GpuType=T4, Gpu=100, RealGpu=25.
	// 例2 实例的机型带有4张gpu整卡, 每张卡对应1张实际T4卡, 则 此时 GpuType=T4, Gpu=400, RealGpu=400.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`

	// Gpu卡型号 T4或者V100
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuType *string `json:"GpuType,omitempty" name:"GpuType"`

	// 创建或更新时无需填写，仅展示需要关注
	// 后付费非整卡实例对应的实际的Gpu卡资源, 表示gpu资源对应实际的gpu卡个数.
	// RealGpu=100表示实际使用了一张gpu卡, 对应实际的实例机型, 有可能代表带有1/4卡的实例4个, 或者带有1/2卡的实例2个, 或者带有1卡的实力1个.
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealGpu *uint64 `json:"RealGpu,omitempty" name:"RealGpu"`
}

type RowItem struct {
	// rowValue 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*RowValue `json:"Values,omitempty" name:"Values"`
}

type RowValue struct {
	// 列名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 列值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SchemaInfo struct {
	// 长度30字符内
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SegmentationInfo struct {
	// 点坐标数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Points []*PointInfo `json:"Points,omitempty" name:"Points"`

	// 分割标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 灰度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gray *uint64 `json:"Gray,omitempty" name:"Gray"`

	// 颜色
	// 注意：此字段可能返回 null，表示取不到有效值。
	Color *string `json:"Color,omitempty" name:"Color"`
}

type SpecPrice struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 原价，单位：分。最大值42亿，超过则返回0
	TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 优惠后的价格，单位：分
	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type SpecUnit struct {
	// 计费项名称
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 计费项数量,建议不超过100万
	SpecCount *uint64 `json:"SpecCount,omitempty" name:"SpecCount"`
}

type StartCmdInfo struct {
	// 启动命令
	StartCmd *string `json:"StartCmd,omitempty" name:"StartCmd"`

	// ps启动命令
	PsStartCmd *string `json:"PsStartCmd,omitempty" name:"PsStartCmd"`

	// worker启动命令
	WorkerStartCmd *string `json:"WorkerStartCmd,omitempty" name:"WorkerStartCmd"`
}

// Predefined struct for user
type StartTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type StartTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *StartTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartTrainingTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartTrainingTaskResponseParams `json:"Response"`
}

func (r *StartTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskRequestParams struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type StopTrainingTaskRequest struct {
	*tchttp.BaseRequest
	
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *StopTrainingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopTrainingTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopTrainingTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopTrainingTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopTrainingTaskResponseParams `json:"Response"`
}

func (r *StopTrainingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopTrainingTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 多个标签值
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`
}

type TrainingDataPoint struct {

}

type TrainingMetric struct {
	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 数据值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*TrainingDataPoint `json:"Values,omitempty" name:"Values"`

	// 上报的Epoch. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Epochs []*TrainingDataPoint `json:"Epochs,omitempty" name:"Epochs"`

	// 上报的Step. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*TrainingDataPoint `json:"Steps,omitempty" name:"Steps"`

	// 上报的TotalSteps. 可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSteps []*TrainingDataPoint `json:"TotalSteps,omitempty" name:"TotalSteps"`
}

type TrainingModelDTO struct {
	// 模型id
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 模型名称
	TrainingModelName *string `json:"TrainingModelName,omitempty" name:"TrainingModelName"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 模型创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type TrainingModelVersionDTO struct {
	// 模型id
	TrainingModelId *string `json:"TrainingModelId,omitempty" name:"TrainingModelId"`

	// 模型版本id
	TrainingModelVersionId *string `json:"TrainingModelVersionId,omitempty" name:"TrainingModelVersionId"`

	// 模型版本
	TrainingModelVersion *string `json:"TrainingModelVersion,omitempty" name:"TrainingModelVersion"`

	// 模型来源
	TrainingModelSource *string `json:"TrainingModelSource,omitempty" name:"TrainingModelSource"`

	// 创建时间
	TrainingModelCreateTime *string `json:"TrainingModelCreateTime,omitempty" name:"TrainingModelCreateTime"`

	// 创建人uin
	TrainingModelCreator *string `json:"TrainingModelCreator,omitempty" name:"TrainingModelCreator"`

	// 算法框架
	AlgorithmFramework *string `json:"AlgorithmFramework,omitempty" name:"AlgorithmFramework"`

	// 推理环境
	ReasoningEnvironment *string `json:"ReasoningEnvironment,omitempty" name:"ReasoningEnvironment"`

	// 推理环境来源
	ReasoningEnvironmentSource *string `json:"ReasoningEnvironmentSource,omitempty" name:"ReasoningEnvironmentSource"`

	// 模型指标
	TrainingModelIndex *string `json:"TrainingModelIndex,omitempty" name:"TrainingModelIndex"`

	// 训练任务名称
	TrainingJobName *string `json:"TrainingJobName,omitempty" name:"TrainingJobName"`

	// 模型cos路径
	TrainingModelCosPath *CosPathInfo `json:"TrainingModelCosPath,omitempty" name:"TrainingModelCosPath"`

	// 模型名称
	TrainingModelName *string `json:"TrainingModelName,omitempty" name:"TrainingModelName"`

	// 训练任务id
	TrainingJobId *string `json:"TrainingJobId,omitempty" name:"TrainingJobId"`

	// 自定义推理环境
	ReasoningImageInfo *ImageInfo `json:"ReasoningImageInfo,omitempty" name:"ReasoningImageInfo"`

	// 模型版本创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模型处理状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelStatus *string `json:"TrainingModelStatus,omitempty" name:"TrainingModelStatus"`

	// 模型处理进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelProgress *uint64 `json:"TrainingModelProgress,omitempty" name:"TrainingModelProgress"`

	// 模型错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelErrorMsg *string `json:"TrainingModelErrorMsg,omitempty" name:"TrainingModelErrorMsg"`

	// 模型格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingModelFormat *string `json:"TrainingModelFormat,omitempty" name:"TrainingModelFormat"`
}

type TrainingTaskDetail struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 训练任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 主账号uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 子账号uin
	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 训练框架名称，eg：SPARK、TENSORFLOW、PYTORCH、LIGHT
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitempty" name:"FrameworkName"`

	// 训练框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" name:"FrameworkVersion"`

	// 训练模式，eg：PS_WORKER、DDP、MPI、HOROVOD
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitempty" name:"TrainingMode"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 预付费专用资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 资源配置
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitempty" name:"ResourceConfigInfos"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 自定义镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`

	// 代码包
	CodePackagePath *CosPathInfo `json:"CodePackagePath,omitempty" name:"CodePackagePath"`

	// 启动命令信息
	StartCmdInfo *StartCmdInfo `json:"StartCmdInfo,omitempty" name:"StartCmdInfo"`

	// 数据来源，eg：DATASET、COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 数据配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataConfigs []*DataConfig `json:"DataConfigs,omitempty" name:"DataConfigs"`

	// 调优参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TuningParameters *string `json:"TuningParameters,omitempty" name:"TuningParameters"`

	// 训练输出
	Output *CosPathInfo `json:"Output,omitempty" name:"Output"`

	// 是否上报日志
	LogEnable *bool `json:"LogEnable,omitempty" name:"LogEnable"`

	// 日志配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogConfig *LogConfig `json:"LogConfig,omitempty" name:"LogConfig"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 训练开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中
	ChargeStatus *string `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 最近一次实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestInstanceId *string `json:"LatestInstanceId,omitempty" name:"LatestInstanceId"`

	// TensorBoard ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TensorBoardId *string `json:"TensorBoardId,omitempty" name:"TensorBoardId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 训练结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 计费金额信息，eg：2.00元/小时 (for后付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInfo *string `json:"BillingInfo,omitempty" name:"BillingInfo"`

	// 预付费专用资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`
}

type TrainingTaskSetItem struct {
	// 训练任务ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 训练任务名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 框架名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkName *string `json:"FrameworkName,omitempty" name:"FrameworkName"`

	// 训练框架版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameworkVersion *string `json:"FrameworkVersion,omitempty" name:"FrameworkVersion"`

	// 训练模式eg：PS_WORKER、DDP、MPI、HOROVOD
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrainingMode *string `json:"TrainingMode,omitempty" name:"TrainingMode"`

	// 计费模式
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 计费状态，eg：BILLING计费中，ARREARS_STOP欠费停止，NOT_BILLING不在计费中
	ChargeStatus *string `json:"ChargeStatus,omitempty" name:"ChargeStatus"`

	// 预付费专用资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" name:"ResourceGroupId"`

	// 资源配置
	ResourceConfigInfos []*ResourceConfigInfo `json:"ResourceConfigInfos,omitempty" name:"ResourceConfigInfos"`

	// 标签配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeInSeconds *uint64 `json:"RuntimeInSeconds,omitempty" name:"RuntimeInSeconds"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 训练开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 训练结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 训练输出
	Output *CosPathInfo `json:"Output,omitempty" name:"Output"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 计费金额信息，eg：2.00元/小时 (for后付费)
	BillingInfo *string `json:"BillingInfo,omitempty" name:"BillingInfo"`

	// 预付费专用资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" name:"ResourceGroupName"`

	// 自定义镜像信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageInfo *ImageInfo `json:"ImageInfo,omitempty" name:"ImageInfo"`
}