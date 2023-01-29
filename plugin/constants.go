package plugin

const (
	// Run command hooks.
	OnConfigLoaded string = "onConfigLoaded"
	OnNewLogger    string = "onNewLogger"
	OnNewPool      string = "onNewPool"
	OnNewClient    string = "onNewClient"
	OnNewProxy     string = "onNewProxy"
	OnNewServer    string = "onNewServer"
	OnSignal       string = "onSignal"
	// Server hooks.
	OnRun      string = "onRun"
	OnBooting  string = "onBooting"
	OnBooted   string = "onBooted"
	OnOpening  string = "onOpening"
	OnOpened   string = "onOpened"
	OnClosing  string = "onClosing"
	OnClosed   string = "onClosed"
	OnTraffic  string = "onTraffic"
	OnShutdown string = "onShutdown"
	OnTick     string = "onTick"
	// Proxy hooks.
	OnTrafficFromClient string = "onTrafficFromClient"
	OnTrafficToServer   string = "onTrafficToServer"
	OnTrafficFromServer string = "onTrafficFromServer"
	OnTrafficToClient   string = "onTrafficToClient"
)
