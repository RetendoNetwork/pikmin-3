package nex

import (
	"github.com/PretendoNetwork/nex-protocols-go/ranking"
	"github.com/RetendoNetwork/pikmin-3/globals"
	nex_ranking "github.com/RetendoNetwork/pikmin-3/nex/ranking"
)

func registerSecureServerNEXProtocols() {
	rankingProtocol := ranking.NewRankingProtocol(globals.SecureServer)

	rankingProtocol.UploadScore(nex_ranking.UploadScore)
	rankingProtocol.GetRanking(nex_ranking.GetRanking)
}
