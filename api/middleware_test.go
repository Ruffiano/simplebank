package api

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct{
		name string
		setupAuth func (*testing.T, request *http.Request, tokenMaker token.Maker)  {
			checkResponse func (t *testing.T, recorder *httptest.ResposeRecorder)
		}{}

		for i := range testCases {
			tc := testCases[i]

			t.Run(tc.name, func (t, *testing.T)  {
				server := newTestServer(t, nil)

				authPath := "/auth"
				server.router.Get(
					authPath,
					authMiddleware(server.tokenMaker),
					func(ctx *gin.Context)  {
						ctx.JSON(http.StatusOK, gin.H{})
					}
				)
			})
		}
	}
}