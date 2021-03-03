# $FreeBSD$

PORTNAME=	r7rs
DISTVERSION=	1.0.3
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Support for most of the R7RS Scheme language

MY_DEPENDS=	${C5_REPO}/matchable.egg-info:devel/chicken5-matchable \
		${C5_REPO}/srfi-1.egg-info:devel/chicken5-srfi-1 \
		${C5_REPO}/srfi-13.egg-info:devel/chicken5-srfi-13
RUN_DEPENDS=	chicken5:lang/chicken5 \
		${MY_DEPENDS}
BUILD_DEPENDS=	chicken-install5:lang/chicken5 \
		${MY_DEPENDS}

CSC_OPTIONS=	-v -strip
C5_REPO:=	${:!chicken-install5 -repository!}
CHICKEN_INSTALL_REPOSITORY=	${STAGEDIR}${C5_REPO}

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && chicken-install5 -v -no-install-dependencies


.export CSC_OPTIONS CHICKEN_INSTALL_REPOSITORY
.include <bsd.port.mk>
