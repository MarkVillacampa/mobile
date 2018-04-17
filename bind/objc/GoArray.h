#import <Foundation/Foundation.h>
#import "ref.h"

@interface GoArray<ObjectType> : NSMutableArray<ObjectType> <goSeqRefInterface>
	@property(strong, readonly) id _ref;
	- (instancetype)initWithRef:(id)ref class:(Class)klass;
	- (instancetype)initWithClass:(Class)klass;
	- (ObjectType)objectAtIndex: (NSUInteger)index;
	- (void)insertObject:(ObjectType)anObject atIndex:(NSUInteger)index;
	- (void)removeObjectAtIndex:(NSUInteger)index;
	- (void)addObject:(ObjectType)anObject;
	- (void)removeLastObject;
	- (void)replaceObjectAtIndex:(NSUInteger)index withObject:(ObjectType)anObject;
@end
