package math

import (
	"database/sql"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/sethvargo/go-password/password"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPassword(t *testing.T) {
	dbPassword, _ := password.Generate(20, 5, 1, false, false)
	assert.NotNil(t, dbPassword)
}

func TestSQLDatabase(t *testing.T) {
	var tfFilePath = "../storage"
	t.Parallel()
	uniquePostfix := strings.ToLower(random.UniqueId())
	expectedUser := "pgadmin"
	expectedPassword, _ := password.Generate(20, 5, 1, false, false)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tfFilePath,
		Vars: map[string]interface{}{
			"postfix":     uniquePostfix,
			"db_user":     expectedUser,
			"db_password": expectedPassword,
		},
		NoColor: false,
	})
	subscriptionID := "03c02aa0-4749-45c2-b078-2cc2bf629317"

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	expectedServername := "postgresqlserver-" + uniquePostfix // see fixture
	actualServername := terraform.Output(t, terraformOptions, "servername")
	rgName := terraform.Output(t, terraformOptions, "rgname")
	expectedSkuName := terraform.Output(t, terraformOptions, "sku_name")
	actualServer := azure.GetPostgreSQLServer(t, rgName, actualServername, subscriptionID)
	actualServerAddress := *actualServer.ServerProperties.FullyQualifiedDomainName
	actualServerUser := *actualServer.ServerProperties.AdministratorLogin

	// Expectation
	assert.NotNil(t, actualServer)
	assert.Equal(t, expectedUser, actualServerUser)
	assert.Equal(t, expectedServername, actualServername)
	assert.Equal(t, expectedSkuName, *actualServer.Sku.Name)

	//Connect Server By DB Client
	ConnectDB(t, expectedUser, expectedPassword, actualServerAddress, actualServername)

}

func ConnectDB(t *testing.T, userName string, expectedPassword string, databaseAddress string, actualServername string) {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", databaseAddress, userName+"@"+actualServername, expectedPassword, "postgres")
	print(connectionString)
	db, err := sql.Open("postgres", connectionString)
	assert.Nil(t, err, "open db failed")
	err = db.Ping()
	assert.Nil(t, err, "connect db failed")
	fmt.Println("Successfully created connection to database")
	var currentTime string
	err = db.QueryRow("select now()").Scan(&currentTime)
	assert.Nil(t, err, "query failed ")
	assert.NotEmpty(t, currentTime, "Get Query Time "+currentTime)

}
