package main

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/forbole/juno/v5/cmd"
	initcmd "github.com/forbole/juno/v5/cmd/init"
	parsetypes "github.com/forbole/juno/v5/cmd/parse/types"
	startcmd "github.com/forbole/juno/v5/cmd/start"
	"github.com/forbole/juno/v5/modules/messages"

	migratecmd "github.com/forbole/bdjuno/v4/cmd/migrate"
	parsecmd "github.com/forbole/bdjuno/v4/cmd/parse"

	"github.com/forbole/bdjuno/v4/types/config"

	// _ "cosmossdk.io/simapp"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v4/database"
	"github.com/forbole/bdjuno/v4/modules"
	"github.com/sourcenetwork/sourcehub/app"
)

func main() {
	initCfg := initcmd.NewConfig().
		WithConfigCreator(config.Creator)

	parseCfg := parsetypes.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(config.MakeEncodingConfig(getBasicManagers())).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("bdjuno").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.NewInitCmd(cfg.GetInitConfig()),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(cfg.GetName(), cfg.GetParseConfig()),
		startcmd.NewStartCmd(cfg.GetParseConfig()),
	)

	executor := cmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// getBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func getBasicManagers() []module.BasicManager {
	// initialize sdk config to use `source` as account prefix - used by bank module
	sdkCfg := sdk.GetConfig()
	sdkCfg.SetBech32PrefixForAccount(app.AccountAddressPrefix, "")

	var appBuilder *runtime.AppBuilder

	config := depinject.Configs(
		app.AppConfig(),
		depinject.Supply(
			log.NewNopLogger(),
		),
	)

	err := depinject.Inject(config, &appBuilder)
	if err != nil {
		panic(err)
	}

	return []module.BasicManager{
		runtime.ProvideBasicManager(appBuilder),
	}
}

// getAddressesParser returns the messages parser that should be used to get the users involved in
// a specific message.
// This should be edited by custom implementations if needed.
func getAddressesParser() messages.MessageAddressesParser {
	return messages.JoinMessageParsers(
		messages.CosmosMessageAddressesParser,
	)
}
