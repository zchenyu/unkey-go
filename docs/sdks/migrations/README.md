# Migrations
(*Migrations*)

### Available Operations

* [V1MigrationsCreateKeys](#v1migrationscreatekeys)
* [V1MigrationsEnqueueKeys](#v1migrationsenqueuekeys)

## V1MigrationsCreateKeys

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := []operations.RequestBody{
        operations.RequestBody{
            APIID: "api_123",
            Name: unkeygo.String("my key"),
            Start: unkeygo.String("unkey_32kq"),
            OwnerID: unkeygo.String("team_123"),
            Meta: map[string]any{
                "billingTier": "PRO",
                "trialEnds": "2023-06-16T17:16:37.161Z",
            },
            Roles: []string{
                "admin",
                "finance",
            },
            Permissions: []string{
                "domains.create_record",
                "say_hello",
            },
            Expires: unkeygo.Int64(1623869797161),
            Remaining: unkeygo.Int64(1000),
            Refill: &operations.V1MigrationsCreateKeysRefill{
                Interval: operations.V1MigrationsCreateKeysIntervalDaily,
                Amount: 100,
            },
            Ratelimit: &operations.V1MigrationsCreateKeysRatelimit{
                Limit: 10,
                RefillRate: 1,
                RefillInterval: 60,
            },
            Enabled: unkeygo.Bool(false),
        },
    }
    ctx := context.Background()
    res, err := s.Migrations.V1MigrationsCreateKeys(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]operations.RequestBody](../../.md)                    | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.V1MigrationsCreateKeysResponse](../../models/operations/v1migrationscreatekeysresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## V1MigrationsEnqueueKeys

### Example Usage

```go
package main

import(
	"os"
	unkeygo "github.com/unkeyed/unkey-go"
	"github.com/unkeyed/unkey-go/models/operations"
	"context"
	"log"
)

func main() {
    s := unkeygo.New(
        unkeygo.WithSecurity(os.Getenv("BEARER_AUTH")),
    )
    request := operations.V1MigrationsEnqueueKeysRequestBody{
        MigrationID: "<value>",
        APIID: "<value>",
        Keys: []operations.Keys{
            operations.Keys{
                Name: unkeygo.String("my key"),
                Start: unkeygo.String("unkey_32kq"),
                OwnerID: unkeygo.String("team_123"),
                Meta: map[string]any{
                    "billingTier": "PRO",
                    "trialEnds": "2023-06-16T17:16:37.161Z",
                },
                Roles: []string{
                    "admin",
                    "finance",
                },
                Permissions: []string{
                    "domains.create_record",
                    "say_hello",
                },
                Expires: unkeygo.Int64(1623869797161),
                Remaining: unkeygo.Int64(1000),
                Refill: &operations.V1MigrationsEnqueueKeysRefill{
                    Interval: operations.V1MigrationsEnqueueKeysIntervalDaily,
                    Amount: 100,
                },
                Ratelimit: &operations.V1MigrationsEnqueueKeysRatelimit{
                    Limit: 10,
                    Duration: 60000,
                },
                Enabled: unkeygo.Bool(false),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Migrations.V1MigrationsEnqueueKeys(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.V1MigrationsEnqueueKeysRequestBody](../../models/operations/v1migrationsenqueuekeysrequestbody.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |


### Response

**[*operations.V1MigrationsEnqueueKeysResponse](../../models/operations/v1migrationsenqueuekeysresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.ErrBadRequest          | 400                              | application/json                 |
| sdkerrors.ErrUnauthorized        | 401                              | application/json                 |
| sdkerrors.ErrForbidden           | 403                              | application/json                 |
| sdkerrors.ErrNotFound            | 404                              | application/json                 |
| sdkerrors.ErrConflict            | 409                              | application/json                 |
| sdkerrors.ErrTooManyRequests     | 429                              | application/json                 |
| sdkerrors.ErrInternalServerError | 500                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |
