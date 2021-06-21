package main

//InitiatorConfig represents initiator's configuration
type InitiatorConfig struct {
	Other  string
	Kind   string
	Name   string
	Target string
}

//TriggerConfig represents trigger's configuration
type TriggerConfig struct {
	Other  string
	Kind   string
	Name   string
	Target string
}

//ConnectConfig contains the extend info for one stage
type ConnectConfig struct {
	Initiator InitiatorConfig
	Trigger   TriggerConfig
}

//RoamingConfig contains the roaming info
type RoamingConfig struct {
	Connects map[string]ConnectConfig
}
