/*�����дһ��������������һ���������� [f1, f2, f3������ fn] ��������һ���µĺ��� fn �����Ǻ�������� ���Ϻ��� ��

[f(x)�� g(x)�� h(x)] �� ���Ϻ��� Ϊ fn(x) = f(g(h(x))) ��

һ���պ����б�� ���Ϻ��� �� ��Ⱥ��� f(x) = x ��

����Լ��������е�ÿ����������һ�����Ͳ�����Ϊ���룬������һ��������Ϊ�����*/ 


var compose = function(functions) {
    
	return function(x) {
        return functions.reduceRight((acc,fn),x=>fn(acc),x)
    }
};

/**
 * const fn = compose([x => x + 1, x => 2 * x])
 * fn(4) // 9
 */