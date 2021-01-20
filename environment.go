package core

const (
	// EnvironmentDev 开发
	EnvironmentDev Environment = "dev"
	// EnvironmentQa 测试
	EnvironmentQa Environment = "qa"
	// EnvironmentProd 生产
	EnvironmentProd Environment = "prod"
	// EnvironmentLocal 本地环境
	EnvironmentLocal Environment = "local"
	// EnvironmentSimulation 模拟请求（不发真实请求到服务器）
	EnvironmentSimulation Environment = "simulation"
)

// Environment 环境类型
type Environment string
