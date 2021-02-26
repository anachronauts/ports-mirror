# $FreeBSD$

PORTNAME=	gemcert
DISTVERSION=	g20200801
PORTREVISION=	0
CATEGORIES=	security

MAINTAINER=	sage@anachronauts.club
COMMENT=	A simple tool for creating self-signed certs for use in Geminispace

LICENSE=	BSD2CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

MASTER_SITES=	https://tildegit.org/solderpunk/gemcert/archive/
DISTFILES=	fc14deb275.tar.gz
WRKSRC=		${WRKDIR}/${PORTNAME}

USES=		go
GO_PKGNAME=	tildegit.org/solderpunk/gemcert

#USE_GITHUB=	nodefault
#GH_TUPLE=	BurntSushi:toml:ea60c4def9:BurntSushi_toml/vendor/github.com/BurntSushi/toml

.include <bsd.port.mk>
