// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build jvm

#include <jni.h>
#include "_cgo_export.h"
#include <stdlib.h>

JNIEXPORT void JNICALL
Java_go_Seq_setContext(JNIEnv* env, jclass clazz, jobject ctx) {
	JavaVM* vm;
	if ((*env)->GetJavaVM(env, &vm) != 0) {
		printf("failed to get JavaVM");
		abort();
	}
	setContext(vm, (*env)->NewGlobalRef(env, ctx));
}
