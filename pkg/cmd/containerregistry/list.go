// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package containerregistry

import (
	"context"

	"github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/internal/util/apiclient/server"
	"github.com/daytonaio/daytona/pkg/cmd/output"
	"github.com/daytonaio/daytona/pkg/views"
	containerregistry_view "github.com/daytonaio/daytona/pkg/views/containerregistry/list"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var containerRegistryListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists container registries",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		apiClient, err := server.GetApiClient(nil)
		if err != nil {
			log.Fatal(err)
		}

		containerRegistries, res, err := apiClient.ContainerRegistryAPI.ListContainerRegistries(context.Background()).Execute()
		if err != nil {
			log.Fatal(apiclient.HandleErrorResponse(res, err))
		}

		if len(containerRegistries) == 0 {
			views.RenderInfoMessage("No container registries found. Set a new container registry by running 'daytona container-registry set'")
			return
		}

		if output.FormatFlag != "" {
			output.Output = containerRegistries
			return
		}

		err = containerregistry_view.ListRegistries(containerRegistries)
		if err != nil {
			log.Fatal(err)
		}
	},
}
