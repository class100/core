package core

const (
	// EnvironmentTypeDev 开发
	EnvironmentTypeDev Environment = "dev"
	// EnvironmentTypeTest 测试
	EnvironmentTypeTest Environment = "test"
	// EnvironmentTypeQa 测试
	EnvironmentTypeQa Environment = "qa"
	// EnvironmentTypeProd 生产
	EnvironmentTypeProd Environment = "prod"
	// EnvironmentTypeLocal 本地环境
	EnvironmentTypeLocal Environment = "local"
	// EnvironmentTypeSimulation 模拟请求（不发真实请求到服务器）
	EnvironmentTypeSimulation Environment = "simulation"
)

// Environment 环境类型
type Environment string
