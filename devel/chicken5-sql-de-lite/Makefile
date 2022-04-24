PORTNAME=	sql-de-lite
DISTVERSION=	0.9.0
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Interface to SQLite 3 for CHICKEN

LIB_DEPENDS=	libsqlite3.so:databases/sqlite3
MY_DEPENDS=	${PREFIX}/lib/chicken5/11/foreigners.egg-info:devel/chicken5-foreigners \
		${PREFIX}/lib/chicken5/11/object-evict.egg-info:devel/chicken5-object-evict \
		${PREFIX}/lib/chicken5/11/srfi-1.egg-info:devel/chicken5-srfi-1 \
		${PREFIX}/lib/chicken5/11/srfi-18.egg-info:devel/chicken5-srfi-18 \
		${PREFIX}/lib/chicken5/11/srfi-69.egg-info:devel/chicken5-srfi-69
RUN_DEPENDS=	chicken5:lang/chicken5 \
		${MY_DEPENDS}
BUILD_DEPENDS=	chicken-install5:lang/chicken5 \
		${MY_DEPENDS}

CSC_OPTIONS=	-v -strip -I${PREFIX}/include -L${PREFIX}/lib

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-D sql-de-lite-external-lib \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && DESTDIR=${STAGEDIR} chicken-install5 \
		-D sql-de-lite-external-lib \
		-v -no-install-dependencies

.export CSC_OPTIONS
.include <bsd.port.mk>
