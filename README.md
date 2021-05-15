# sr_auth

## Example

```
key := "SECRETKEY"

auth := sr_auth.CreateAuth(key)
user, err := auth.GetUserFromToken(tokenString)
if err != nil {
    log.Fatal(err)
}
```