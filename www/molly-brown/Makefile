# $FreeBSD$

PORTNAME=	molly-brown
DISTVERSION=	g20210124
PORTREVISION=	1
CATEGORIES=	www

MAINTAINER=	sage@anachronauts.club
COMMENT=	a full-featured Gemini server implemented in Go

LICENSE=	BSD2CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

MASTER_SITES=	https://tildegit.org/solderpunk/molly-brown/archive/
DISTFILES=	92cd40db12.tar.gz
WRKSRC=		${WRKDIR}/${PORTNAME}

USES=		go
GO_PKGNAME=	tildegit.org/solderpunk/molly-brown

USE_GITHUB=	nodefault
GH_TUPLE=	BurntSushi:toml:ea60c4def9:BurntSushi_toml/vendor/github.com/BurntSushi/toml

USERS=		gemini
GROUPS=		gemini

USE_RC_SUBR=	molly_brown

post-patch:
	@${REINPLACE_CMD} -e 's|/etc/molly.conf|${PREFIX}/etc/molly.conf|g' ${WRKSRC}/main.go

post-install:
	${INSTALL_DATA} ${WRKSRC}/example.conf ${STAGEDIR}${PREFIX}/etc/molly.conf.sample

.include <bsd.port.mk>
