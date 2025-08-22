package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_region "github.com/apimanagementclient/mcp-server/tools/region"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_region.CreateRegion_listbyserviceTool(cfg),
	}
}
