# $FreeBSD$

PORTNAME=	gemlikes
DISTVERSIONPREFIX=	v
DISTVERSION=	1.0.2-7
DISTVERSIONSUFFIX=	-g0bbe9b1
CATEGORIES=	www

MAINTAINER=	sage@anachronauts.club
COMMENT=	Liking and comment system for the Gemini protocol

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	makeworld-the-better-one
GH_TUPLE=	icza:backscanner:2120fccb01f7:icza_backscanner/vendor/github.com/icza/backscanner \
		pelletier:go-toml:v1.8.0:pelletier_go_toml/vendor/github.com/pelletier/go-toml

GO_TARGET=	./add-comment:${WWWDIR}/add-comment \
		./like:${WWWDIR}/like \
		./view:${WWWDIR}/view

SUB_FILES=	gemlikes-install

post-install:
	${INSTALL_DATA} ${WRKSRC}/example-config.toml ${STAGEDIR}${WWWDIR}/gemlikes.toml.sample
	${INSTALL_SCRIPT} ${WRKDIR}/gemlikes-install ${STAGEDIR}${PREFIX}/bin

.include <bsd.port.mk>
