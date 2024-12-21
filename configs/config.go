package configs 

imports (
	"log"
	"os"
	"sync"
	
	"github.com./spf13/viper"
	"github.com/joho/godotenv"
)

// config holds the application configuration

type config struct {
	AppName string `mapstructure:"app_name"`
	ServerPort string `mapstructure:"server_port"`
	DatabaseURL string `mapstructure:"database_url"`
	AuthService string `mapstructure: "auth_service"`
	WalletService string `mapstructure:"wallet_service`
	TransactionService string `mapstructure:transaction_service`
	JWTSecret 		string 	`mapstructure:"jwt_secret"`
	Environment		string	`mapstructure:"environment"`
}

var (
	config *config
	once  sync.Once
)

// Loadconfig initializes the configuration from app_config.yaml
func LoadConfig () *Config { 
	once.Do(func(){
		// Load environment variables from secrets 
		if err : = godotenv.Load("configs/secrets.env "); err !=nil {
			
		}
		
		//Load YAML configuration using viper 
		viper.SetConfigName("app_config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("configs/")
		
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v ", err)
		}
		
		//Unmarshal config into struct
		config = &Config{}
		if err := viper.Unmarshall(config); err != nil {
			log.Fatalf("Unable to marshal config: %v", err)
		}
		
		
			//Override config values with environment variablles, if 
		config.JWTSecret = os.Getenv("JWT_SECRET")
	})
	
	return config
}