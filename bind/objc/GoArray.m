#import "GoArray.h"
#include "_cgo_export.h"

@implementation GoArray 
{
	Class _klass;
}
- (instancetype)initWithRef:(id)ref class:(Class)klass {
	self = [super init];
	if (self) { 
	    __ref = ref; 
	    _klass = klass;
	}
	return self;
}

- (instancetype)initWithClass:(Class)klass {

	self = [super init];
	if (self) {
		_klass = klass;
		GoSeqRef* refnum = [[_klass new] _ref]; 
		__ref = go_seq_from_refnum(go_new_slice(go_seq_go_to_refnum(refnum)));
	}
	return self;
}

- (NSUInteger)count
{
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    return go_count(refnum);
}

- (id)objectAtIndex: (NSUInteger)index
{
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    int32_t r0 = go_objectAtIndex(refnum, index);
    NSLog(@"refnum: %d", r0);
    id _r0 = nil;
    GoSeqRef* _r0_ref = go_seq_from_refnum(r0);
    if (_r0_ref != NULL) {
	_r0 = _r0_ref.obj;
	if (_r0 == nil) {
		_r0 = [[_klass alloc] initWithRef:_r0_ref];
	}
    }
    return _r0;

/*     int32_t refnum = go_seq_go_to_refnum(self._ref); */
/*     return [NSNumber numberWithInt: go_objectAtIndex(refnum, index)]; */
}

- (void)insertObject:(id)obj atIndex:(NSUInteger)index 
{
    int32_t objref;
    if ([obj conformsToProtocol:@protocol(goSeqRefInterface)]) {
	    id<goSeqRefInterface> v_proxy = (id<goSeqRefInterface>)(obj);
	    objref = go_seq_go_to_refnum(v_proxy._ref);
    } else {
	    objref = go_seq_to_refnum(obj);
    }
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    go_insertObjectAtIndex(refnum, objref, index);
}

- (void)removeObjectAtIndex:(NSUInteger)index {
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    go_removeObjectAtIndex(refnum, index);
}

- (void)addObject:(id)obj {
    int32_t objref;
    if ([obj conformsToProtocol:@protocol(goSeqRefInterface)]) {
	    id<goSeqRefInterface> v_proxy = (id<goSeqRefInterface>)(obj);
	    objref = go_seq_go_to_refnum(v_proxy._ref);
    } else {
	    objref = go_seq_to_refnum(obj);
    }
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    go_addObject(refnum, objref);
}

- (void)removeLastObject {
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    go_removeLastObject(refnum);
}

- (void)replaceObjectAtIndex:(NSUInteger)index withObject:(id)obj {
    int32_t objref;
    if ([obj conformsToProtocol:@protocol(goSeqRefInterface)]) {
	    id<goSeqRefInterface> v_proxy = (id<goSeqRefInterface>)(obj);
	    objref = go_seq_go_to_refnum(v_proxy._ref);
    } else {
	    objref = go_seq_to_refnum(obj);
    }
    int32_t refnum = go_seq_go_to_refnum(self._ref);
    go_replaceObjectAtIndexWithObject(refnum, index, objref);
}
@end
