# GCF Token Auth

Simple helper (insecure) to auth with a token in header when calling Google Cloud function and compare it with environment variable named `X_PROXIMIS_API_KEY`.

## Usage

Add "X_PROXIMIS_API_KEY" environment variable in your Google Cloud Function.

Call your function with token header :

```bash
curl https://ZONE-PROJECT.cloudfunctions.net/YOURFUNCTION -H "X-Proximis-Api-Key=YOURMARVELLOUSTOKEN"
```

If token is different from "X_PROXIMIS_API_KEY" environment variable it will launch http error and Golang panic func.

```golang
package p

import (
    gcftokenauth "github.com/proximisfr/gcf-token-auth"
)

func F(w http.ResponseWriter, r *http.Request) {
    if !gcftokenauth.Auth(w, r) {
        return
    }
    [...]
}
```
