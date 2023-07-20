package HTTPHelper

type HTTPCodes struct {
	UNAUTHORIZED    int    `default:401`
	UNAUTHENTICATED int    `default:403`
	SUCCESS         int    `default:200`
	NOTSUCCESS      int    `default:404`
	GET             string `default:"GET"`
	POST            string `default:"POST"`
	PUT             string `default:"PUT"`
	DELETE          string `default:"DEL"`
}
