# $FreeBSD$

PORTNAME=	sql-de-lite
DISTVERSION=	0.8.1
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Interface to SQLite 3 for CHICKEN

LIB_DEPENDS=	libsqlite3.so:databases/sqlite3
MY_DEPENDS=	${C5_REPO}/foreigners.egg-info:devel/chicken5-foreigners \
		${C5_REPO}/object-evict.egg-info:devel/chicken5-object-evict \
		${C5_REPO}/srfi-18.egg-info:devel/chicken5-srfi-18 \
		${C5_REPO}/srfi-69.egg-info:devel/chicken5-srfi-69
RUN_DEPENDS=	chicken5:lang/chicken5 \
		${MY_DEPENDS}
BUILD_DEPENDS=	chicken-install5:lang/chicken5 \
		${MY_DEPENDS}

CSC_OPTIONS=	-v -strip -I${PREFIX}/include -L${PREFIX}/lib
C5_REPO:=	${:!chicken-install5 -repository!}

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-D sql-de-lite-external-lib \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && DESTDIR=${STAGEDIR} chicken-install5 \
		-D sql-de-lite-external-lib \
		-v -no-install-dependencies


.export CSC_OPTIONS CHICKEN_INSTALL_REPOSITORY
.include <bsd.port.mk>
