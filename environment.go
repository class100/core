package core

const (
	// EnvironmentTypeDev 开发
	EnvironmentTypeDev EnvironmentType = "dev"
	// EnvironmentTypeTest 测试
	EnvironmentTypeTest EnvironmentType = "test"
	// EnvironmentTypeProd 生产
	EnvironmentTypeProd EnvironmentType = "prod"
	// EnvironmentTypeLocal 本地环境
	EnvironmentTypeLocal EnvironmentType = "local"
	// EnvironmentTypeSimulation 模拟请求（不发真实请求到服务器）
	EnvironmentTypeSimulation EnvironmentType = "simulation"
)

// EnvironmentType 环境类型
type EnvironmentType string
