// Retrieves or creates projects in any given Atlas organization
package project


type ProjectSettings struct {
	IsCollectDatabaseSpecificsStatisticsEnabled *bool `field:"optional" json:"isCollectDatabaseSpecificsStatisticsEnabled" yaml:"isCollectDatabaseSpecificsStatisticsEnabled"`
	IsDataExplorerEnabled *bool `field:"optional" json:"isDataExplorerEnabled" yaml:"isDataExplorerEnabled"`
	IsPerformanceAdvisorEnabled *bool `field:"optional" json:"isPerformanceAdvisorEnabled" yaml:"isPerformanceAdvisorEnabled"`
	IsRealtimePerformancePanelEnabled *bool `field:"optional" json:"isRealtimePerformancePanelEnabled" yaml:"isRealtimePerformancePanelEnabled"`
	IsSchemaAdvisorEnabled *bool `field:"optional" json:"isSchemaAdvisorEnabled" yaml:"isSchemaAdvisorEnabled"`
}

