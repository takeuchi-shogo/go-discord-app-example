package setting

type (
	Infrastructure struct {
		MySQL struct {
			DBMaster Database
			DBSlave  Database
		}
		Redis struct {
			Redis Redis
		}
	}

	Database struct {
		Host     string
		User     string
		Password string
		DBName   string
	}

	Redis struct {
		Host string
		Port string
	}
)
