var myApp = angular.module('recordsApp', ['ui.bootstrap']);


myApp.controller('recordsController', function($scope, $http) {
	$scope.filteredRecords = []
	,$scope.currentPage = 1
	,$scope.numPerPage = 10
	,$scope.maxSize = 5;

	$scope.loading = true;
	$http.get('/records').then(function successCallback(response) {
  		$scope.records = response.data;

		$scope.$watch("currentPage + numPerPage", function() {
	    	var begin = (($scope.currentPage - 1) * $scope.numPerPage), end = begin + $scope.numPerPage;
		    $scope.filteredRecords = $scope.records.slice(begin, end);
		});
		$scope.loading = false;
  	});
	



});
